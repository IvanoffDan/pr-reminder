package service

import (
	"fmt"

	"github.com/IvanoffDan/pr-reminder/models"
	"github.com/IvanoffDan/pr-reminder/utils"
	resty "gopkg.in/resty.v1"
)

// BaseURL to connect to BitBucket
const BaseURL = "https://api.bitbucket.org/2.0"

// BBService is a service to connect to BitBucket
type BBService struct {
	config utils.Config
}

// NewBBService return new instance of BitBucket service
func NewBBService(appConfig utils.Config) Service {
	return &BBService{
		appConfig,
	}
}

// GetPullRequests fetches all pull requests from a repo
func (bbs *BBService) GetPullRequests() ([]models.PullRequest, error) {
	client := bbs.newClient()

	result := &struct {
		Values []models.PullRequest `json:"values"`
	}{}

	req := client.R().SetResult(result)

	_, err := req.Get(fmt.Sprintf(
		"%v/repositories/%v/pullrequests",
		BaseURL,
		bbs.config.GetRepoSlug()))

	if err != nil {
		log.Errorf("Error - BBService - GetPullRequests - %v", err.Error())
		return nil, fmt.Errorf("Could not fetch pull request from %v repository", bbs.config.GetRepoSlug())
	}

	log.Infof("GetPullRequests resp: %+v", result)

	prs := result.Values

	// Fetching reviewers for each Pull Request
	for _, pr := range prs {
		fetchedPR, err := bbs.GetPullRequest(pr.ID)

		if err != nil {
			log.Errorf("Error - BBService - GetPullRequests - %v", err.Error())
			log.Errorf("Could not fetch reviewer for pull request #%v from %v repository", pr.ID, bbs.config.GetRepoSlug())
			// Add an empty slice
			pr.AssignReviewers([]models.Reviewer{})
		}

		pr.AssignReviewers(fetchedPR.Reviewers)
	}

	return prs, nil
}

// GetPullRequest fetches PullRequest by ID
func (bbs *BBService) GetPullRequest(id int) (*models.PullRequest, error) {
	client := bbs.newClient()

	result := &models.PullRequest{}

	req := client.R().SetResult(result)

	resp, err := req.Get(fmt.Sprintf(
		"%v/repositories/%v/pullrequests/%v",
		BaseURL,
		bbs.config.GetRepoSlug(),
		id,
	))

	log.Info(resp.StatusCode())

	if err != nil {
		log.Errorf("Error - BBService - GetPullRequest - %v", err.Error())
		return result, fmt.Errorf("Could not fetch pull request %v from %v repository", id, bbs.config.GetRepoSlug())
	}

	log.Infof("Fetched PR#%v %+v", id, result)

	return result, nil
}

func (bbs *BBService) newClient() *resty.Client {
	client := resty.New()
	client.SetBasicAuth(bbs.config.GetUsername(), bbs.config.GetPassword())

	return client
}
