/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "exciting-opendigger/service"
	"fmt"

	"github.com/spf13/cobra"
)

var position string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download data from api",
	Long:  `download data from api and generate pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")
		// 获取结果
		source = getResult(queryPara)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("download print")
		// var downloadService service.DownloadService
		// downloadService = &service.DownloadAsPdf{"", ""}

		// 打印结果
		fmt.Print(source)
		// downloadService.SetData(source)
		// downloadService.Download(position)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// 下载相关参数
	downloadCmd.Flags().StringVarP(&position, "position", "p", "", "Download place where data to write to")
}
