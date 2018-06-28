package service

import (
	"github.com/IvanoffDan/pr-reminder/models"
	"github.com/IvanoffDan/pr-reminder/utils"
)

// Importing logger in a main file of a package
var log = utils.GetLogger()

// Service to connect to repository
type Service interface {
	GetPullRequests() ([]models.PullRequest, error)
}
