/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"exciting-opendigger/service"
	"fmt"
	"github.com/spf13/cobra"
)

var source string

var position string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download data from api",
	Long:  `download data from api and generate pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")
		var downloadService service.DownloadService
		downloadService = &service.DownloadAsPdf{"", ""}
		downloadService.SetData(source)
		downloadService.Download(position)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&source, "source", "s", "", "Source api to read from")
	downloadCmd.Flags().StringVarP(&position, "position", "p", "", "Download place where data to write to")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
