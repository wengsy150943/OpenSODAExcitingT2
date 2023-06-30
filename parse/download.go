/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"exciting-opendigger/service"
	"fmt"
	"strings"

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
		if (queryPara.metric != "" && queryPara.month != "") {
			panic("Here is not necessary to download the result: simple return value.")
		}
		repoInfo = getResult(queryPara)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		var downloadService service.DownloadService
		fmt.Print(repoInfo)
		downloadService = &service.SingleDownloadService{}

		if (strings.Contains(cmd.CommandPath(), "compare") ){
			// TODO qk: download compare
			downloadService.SetData(repoInfo, position)
		} else{
			downloadService.SetData(repoInfo, position)
		}

		
		downloadService.Download()
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// 下载相关参数
	downloadCmd.Flags().StringVarP(&position, "position", "p", "", "Download place where data to write to")
	downloadCmd.MarkFlagRequired("position")
}
