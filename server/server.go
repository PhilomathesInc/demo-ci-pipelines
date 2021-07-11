package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PhilomathesInc/demo-ci-pipelines/config"
	"go.uber.org/zap"
)

const (
	healthPath  = "/health"
	welcomePath = "/"
)

func Run(c *config.Config) {
	c.Logger.Info("initializing the server...")
	http.HandleFunc(healthPath, healthHandler(c))
	http.HandleFunc(welcomePath, welcomeHandler(c))

	timeout, err := time.ParseDuration(c.Server.Timeout)
	if err != nil {
		c.Logger.Error("error parsing timeout value", zap.String("error", err.Error()))
	}
	s := http.Server{
		Addr:         fmt.Sprintf(":%s", c.Server.Port),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	c.Logger.Info(fmt.Sprintf("you can now visit the server on %s:%s", c.Server.Hostname, c.Server.Port))
	if err := s.ListenAndServe(); err != nil {
		c.Logger.Fatal("error occured on the HTTP server", zap.String("error", err.Error()))
	}
}

func healthHandler(c *config.Config) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		c.Logger.Info("received request to check the server health")
		w.Header().Add("Content-Type", "application/json")
		if _, err := w.Write([]byte(Healthy)); err != nil {
			c.Logger.Error("error writing status response in ResponseWritter", zap.String("error", err.Error()))
		}
	}
}

func welcomeHandler(c *config.Config) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", WelcomeMsg)
	}
}
