package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "demo-cli",
		Short: "A demo Go based CLI application",
		Long:  "The CLI sends request to the demo server application and prints the response received from the server",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Mugdha Adhav", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringP("license", "l", "MIT License", "name of license for the project")
}
