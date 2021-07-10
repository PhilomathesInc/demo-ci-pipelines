package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Demo Application",
	Long:  `Display's the current version of the Demo Application written in Go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Demo App featuring multiple CI pipelines v0.1.0 -- HEAD")
	},
}
