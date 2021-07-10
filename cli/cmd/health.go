package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"PhilomathesInc/demo-ci-pipelines/config"
	"PhilomathesInc/demo-ci-pipelines/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serverResponse string

func init() {
	rootCmd.AddCommand(healthCmd)
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Print the health status of Demo Application",
	Long:  `Checks the health status of the demo application by sending a GET request to the server`,
	Run: func(cmd *cobra.Command, args []string) {
		checkServerStatus()
	},
}

func checkServerStatus() {
	c, err := config.New()
	if err != nil {
		panic(err)
	}

	u := fmt.Sprintf("%s://%s:%s/health", c.Server.Protocol, c.Server.Hostname, c.Server.Port)
	// Syntax check for URL by parsing it into a struct
	if _, err := url.ParseRequestURI(u); err != nil {
		c.Logger.Error("could not parse the server URL", zap.String("error", err.Error()))
		return
	}
	resp, err := http.Get(u)
	if err != nil {
		c.Logger.Error("could not send GET request to the server URL", zap.String("error", err.Error()))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Error("could not read the response from the server", zap.String("error", err.Error()))
		return
	}
	if err := json.Unmarshal(body, &serverResponse); err != nil {
		c.Logger.Error("could not unmarshal the response from the server", zap.String("error", err.Error()))
		return
	}
	if serverResponse == string(server.Healthy) {
		fmt.Println("The server responded with healthy status")
	} else {
		fmt.Println("The server is unhealthy")
	}
	fmt.Println("Response from server: ", serverResponse)
}
