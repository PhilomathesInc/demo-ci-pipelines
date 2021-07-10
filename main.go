package main

import (
	"PhilomathesInc/demo-ci-pipelines/config"
	"PhilomathesInc/demo-ci-pipelines/server"
)

func main() {
	c, err := config.New()
	if err != nil {
		panic(err)
	}
	c.Logger.Info("all configs loaded successfully")

	// Starting the server
	server.Run(c)
}
