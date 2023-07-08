/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"exciting-opendigger/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show previous logs of query",
	Long:  "Show previous logs of query",
	Run: func(cmd *cobra.Command, args []string) {
		var logs []utils.Searchhistory
		utils.Readlog(&logs)
		for it, history := range logs {
			fmt.Printf("%d. %s\n", it + 1, history.Log)
		}	
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
