package main

import (
	"fmt"
	"os"

	"github.com/IvanoffDan/pr-reminder/utils"
	"github.com/sirupsen/logrus"
)

var log = utils.GetLogger()

// Config is an interface to store application configuration extracted from environmental variables
type Config interface {
	//Log logs current app configuration to stdout
	Log()
}

// AppConfig stores application configuration
type AppConfig struct {
	Username       string
	Password       string
	RepositorySlug string
}

//Log logs current app configuration to stdout
func (ac *AppConfig) Log() {
	log.WithFields(logrus.Fields{
		"Username":    ac.Username,
		"AppPassword": fmt.Sprintf("%v characters", len(ac.Password)),
		"Repository":  ac.RepositorySlug,
	}).Info()
}

// NewConfig returns new AppConfig
func NewConfig() Config {
	return &AppConfig{
		getEnvOrExit("BB_USERNAME"),
		getEnvOrExit("BB_APP_PASSWORD"),
		getEnvOrExit("REPO_SLUG"),
	}
}

func getEnvOrExit(varName string) string {
	v := os.Getenv(varName)

	if v == "" {
		log.Fatalf("Environmental variable ${%v} is required but not defined", varName)
	}

	return v
}
