package models

// PullRequest is a pull request fetched from repository
type PullRequest struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	State     string     `json:"state"`
	Reviewers []Reviewer `json:"reviewers"`
}

// AssignReviewers adds reviewers to a PullRequest
func (pr *PullRequest) AssignReviewers(reviewers []Reviewer) {
	pr.Reviewers = reviewers
}
