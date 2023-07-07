/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"exciting-opendigger/service"
	"github.com/spf13/cobra"
	"strings"
)

var position string
var draw bool

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download and plot data from OpenDigger",
	Long:  `Download data from api and generate pdf`,
	PersistentPreRun: func (cmd *cobra.Command, args []string)  {
		isShow = false
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 获取结果
		if queryPara.metric != "" && queryPara.month != "" {
			panic("Here is not necessary to download the result: simple return value.")
		}
		repoInfo = getResult(queryPara)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		

		if strings.Contains(cmd.CommandPath(), "compare") {
			// TODO qk: download compare
			downloadService := &service.CompareDownloadService{}
			downloadService.SetData(repoInfo, repoInfoCompare, position)

			if draw {
				// TODO qk: call plot
			} else {
				downloadService.Download()
			}
		} else {
			downloadService := &service.SingleDownloadService{}
			downloadService.SetData(repoInfo, position)

			if draw {
				// TODO qk: call plot
			} else {
				downloadService.Download()
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// 下载相关参数
	downloadCmd.Flags().StringVarP(&position, "position", "p", "", "Download place where data to write to")
	downloadCmd.MarkFlagRequired("position")
	downloadCmd.Flags().BoolVarP(&draw, "draw", "d", false, "Plot data")
}
