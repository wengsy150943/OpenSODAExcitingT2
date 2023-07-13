/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"exciting-opendigger/service"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var position string
var user bool

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download and plot data from OpenDigger",
	Long:  `Download data from api and generate pdf`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		isShow = false
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 获取结果
		// if queryPara.metric != "" && queryPara.month != "" {
		// 	panic("Here is not necessary to download the result: simple return value.")
		// }
		if queryPara.user != "" {
			userInfo = service.GetCertainUser(queryPara.user)
		} else{
			repoInfo = getResult(queryPara)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {

		if strings.Contains(cmd.CommandPath(), "compare") {

			downloadService := &service.CompareDownloadService{}
			err := downloadService.SetData(repoInfo, repoInfoCompare, position)
			if err != nil {
				panic("fail to SetData")
			}
			err2 := downloadService.Download()
			if err2 != nil {
				panic("fail to Download")
			}
		} else {

			if queryPara.month == "" && queryPara.metric == "" {
				if queryPara.user != "" {
					//TODO:user数据的打印处理
					downloadService := &service.UserDownloadService{}
					err := downloadService.SetData(position, queryPara.user)
					if err != nil {
						panic("fail to SetData")
					}
					err2 := downloadService.Download()
					if err2 != nil {
						panic("fail to Download")
					}
				} else { //全数据下载
					downloadService := &service.SingleDownloadService{}
					err := downloadService.SetData(repoInfo, position)
					if err != nil {
						panic("fail to SetData")
					}
					err2 := downloadService.Download()
					if err2 != nil {
						panic("fail to Download")
					}
				}
			} else if queryPara.month == "" { // 特定指标下载
				downloadService := &service.SingleDownloadService{}
				err := downloadService.SetDataOneMetric(repoInfo, position, queryPara.metric)
				if err != nil {
					panic("fail to SetData")
				}
				err2 := downloadService.Download()
				if err2 != nil {
					panic("fail to Download")
				}

			} else if queryPara.metric == "" { // 特定月份的整体报告下载
				y, err1 := strconv.Atoi(queryPara.month[:4])
				if err1 != nil {
					panic("fail to parse the date's year")
				}
				m, err2 := strconv.Atoi(queryPara.month[5:])
				if err2 != nil {
					panic("fail to parse the date's month")
				}
				downloadService := &service.SingleDownloadService{}
				err3 := downloadService.SetDataOneMonth(repoInfo, position, y, m, "")
				if err3 != nil {
					panic("fail to SetDataOneMonth")
				}
				err4 := downloadService.DownloadMonth()
				if err4 != nil {
					panic("fail to Download DataOneMonth")
				}

			} else {
				y, err1 := strconv.Atoi(queryPara.month[:4])
				if err1 != nil {
					panic("fail to parse the date's year")
				}
				m, err2 := strconv.Atoi(queryPara.month[5:])
				if err2 != nil {
					panic("fail to parse the date's month")
				}
				downloadService := &service.SingleDownloadService{}
				err3 := downloadService.SetDataOneMonth(repoInfo, position, y, m, queryPara.metric)
				if err3 != nil {
					panic("fail to SetDataOneMonth")
				}
				err4 := downloadService.DownloadMonth()
				if err4 != nil {
					panic("fail to Download DataOneMonth")
				}
			}
		}
	},
}

func init() {
	

	// 下载相关参数
	downloadCmd.Flags().StringVarP(&position, "position", "p", "", "Download place where data to write to")
	downloadCmd.MarkFlagRequired("position")
	rootCmd.AddCommand(downloadCmd)
}
