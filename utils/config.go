package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

var log = GetLogger()

// Config is an interface to store application configuration extracted from environmental variables
type Config interface {
	//Log logs current app configuration to stdout
	Log()
	LogUsers()
	GetUsername() string
	GetPassword() string
	GetRepoSlug() string
	ReadUsers()
}

// User represents a user of an organization
type User struct {
	Username string
	Email    string
}

// AppConfig stores application configuration
type AppConfig struct {
	Username       string
	Password       string
	RepositorySlug string
	Users          []User
}

//GetUsername returns username
func (ac *AppConfig) GetUsername() string {
	return ac.Username
}

//GetPassword returns password
func (ac *AppConfig) GetPassword() string {
	return ac.Password
}

//GetRepoSlug returns repository slug
func (ac *AppConfig) GetRepoSlug() string {
	return ac.RepositorySlug
}

//Log logs current app configuration to stdout
func (ac *AppConfig) Log() {
	log.WithFields(logrus.Fields{
		"Username":    ac.Username,
		"AppPassword": fmt.Sprintf("%v characters", len(ac.Password)),
		"Repository":  ac.RepositorySlug,
	}).Info()
}

//LogUsers logs all registered users
func (ac *AppConfig) LogUsers() {
	log.Info("Registered users - Start")
	for _, u := range ac.Users {
		log.WithFields(logrus.Fields{
			"username": u.Username,
			"email":    u.Email,
		}).Info()
	}
	log.Info("Registered users - End")
}

//ReadUsers reads user list from disk
func (ac *AppConfig) ReadUsers() {
	appPath := getEnvOrExit("APP_PATH")

	raw, err := ioutil.ReadFile(fmt.Sprintf("%v/users.json", appPath))

	if err != nil {
		log.Errorf("Error reading users.json file %v", err.Error())
		os.Exit(1)
	}

	var users []User

	err = json.Unmarshal(raw, &users)

	if err != nil {
		log.Errorf("Error reading users.json file %v", err.Error())
		os.Exit(1)
	}

	ac.Users = users
}

// NewConfig returns new AppConfig
func NewConfig() Config {
	appConfig := &AppConfig{
		Username:       getEnvOrExit("BB_USERNAME"),
		Password:       getEnvOrExit("BB_APP_PASSWORD"),
		RepositorySlug: getEnvOrExit("REPO_SLUG"),
	}

	// This will call os.Exit(1) if fails
	appConfig.ReadUsers()

	fmt.Println(appConfig.Users)

	return appConfig
}

func getEnvOrExit(varName string) string {
	v := os.Getenv(varName)

	if v == "" {
		log.Fatalf("Environmental variable ${%v} is required but not defined", varName)
	}

	return v
}
