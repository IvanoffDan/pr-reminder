package main

import (
	"fmt"
)

func main() {
	appConfig := NewConfig()
	appConfig.Log()

	fmt.Println("It works!")
}
