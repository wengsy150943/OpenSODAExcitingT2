/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package parse

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of OpenSODAExcitingT2",
	Long:  `All software has versions. This is OpenSODAExcitingT2's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OpenSODAExcitingT2 version v0.1 beta ")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
