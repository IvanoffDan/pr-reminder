package main

import (
	"fmt"

	"github.com/IvanoffDan/pr-reminder/service"
	"github.com/IvanoffDan/pr-reminder/utils"
)

var log = utils.GetLogger()

func main() {
	appConfig := utils.NewConfig()
	appConfig.Log()

	bbService := service.NewBBService(appConfig)

	prs, _ := bbService.GetPullRequests()

	log.Infof("Fetched %v pull requests", len(prs))

	fmt.Println("It works!")
}
