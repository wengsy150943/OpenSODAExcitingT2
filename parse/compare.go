/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"exciting-opendigger/service"
	"github.com/spf13/cobra"
)

var repoArr []string
var userArr []string
var monthArr []string

// downloadCmd represents the download command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "compare data from api",
	Long:  `compare data from api and print in screen`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(userArr) == 2 {
			userInfo = service.GetCertainUser(userArr[0])
			userInfoCompare = service.GetCertainUser(userArr[1])
			return
		}

		// 参数有效性检验
		if len(repoArr)+len(monthArr) != 3 || len(repoArr)*len(monthArr) == 0 {
			panic("argc number mismatch")
		}

		// 初始化查询，第一个查询总是取数组的第一个元素作为参数，第二个查询有两个参数是取数组的第一个元素，所以先全部用数组的第一个参数初始化
		queryPara1 := Query{
			repo:   repoArr[0],
			month:  monthArr[0],
			metric: queryPara.metric,
		}
		queryPara2 := Query{
			repo:   repoArr[0],
			month:  monthArr[0],
			metric: queryPara.metric,
		}

		// 根据比较的对象修改对应的参数
		if len(repoArr) == 2 {
			queryPara2.repo = repoArr[1]
		} else {
			queryPara2.month = monthArr[1]
		}

		// 获取结果
		repoInfo = getResult(queryPara1)
		repoInfoCompare = getResult(queryPara2)
	},
}

func init() {
	compareCmd.Flags().StringArrayVarP(&repoArr, "repo", "r", nil, "Repositories asked")
	compareCmd.Flags().StringArrayVarP(&userArr, "user", "u", nil, "Users asked")
	compareCmd.MarkFlagsMutuallyExclusive("repo", "user")

	compareCmd.Flags().StringArrayVarP(&monthArr, "month", "M", nil, "Month of asked metric")
	compareCmd.Flags().StringVarP(&queryPara.metric, "metric", "m", "", "Metric asked")

	var showCompareCmd = *compareCmd
	var downloadCompareCmd = *compareCmd
	showCmd.AddCommand(&showCompareCmd)
	downloadCmd.AddCommand(&downloadCompareCmd)
}
