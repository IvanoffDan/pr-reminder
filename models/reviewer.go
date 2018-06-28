package models

// Reviewer represents a user assigned to review a PR
type Reviewer struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}
