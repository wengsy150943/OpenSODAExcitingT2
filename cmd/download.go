/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"exciting-opendigger/service"

	"github.com/spf13/cobra"
)

var position string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download and plot data from OpenDigger",
	Long:  `Download data from api and generate pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取结果
		source = getResult(queryPara)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		var downloadService service.DownloadService

		// 打印结果
		str,_ :=json.Marshal(source)
		downloadService.SetData(string(str), position)
		downloadService.Download()
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// 下载相关参数
	downloadCmd.Flags().StringVarP(&position, "position", "p", "", "Download place where data to write to")
}
