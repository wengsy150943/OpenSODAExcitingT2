/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show previous logs of query",
	Long:  "Show previous logs of query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OpenSODAExcitingT2 version v0.1 beta ")
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
