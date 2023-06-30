/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"
	"github.com/spf13/cobra"
)

var repoArr []string
var monthArr []string

// downloadCmd represents the download command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "compare data from api",
	Long:  `compare data from api and print in screen`,
	Run: func(cmd *cobra.Command, args []string) {
		// 参数有效性检验
		if (len(repoArr) + len(monthArr) != 3 || len(repoArr) * len(monthArr) == 0){
			panic("argc number mismatch")
		}

		// 获取结果
		var ret1 map[string]string
		var ret2 map[string]string

		// 初始化查询，第一个查询总是取数组的第一个元素作为参数，第二个查询有两个参数是取数组的第一个元素，所以先全部用数组的第一个参数初始化
		queryPara1 := Query{
				repo : repoArr[0],
				month : monthArr[0],
				metric : queryPara.metric,
			}
		queryPara2 := Query{
				repo : repoArr[0],
				month : monthArr[0],
				metric : queryPara.metric,
			}

		// 根据比较的对象修改对应的参数
		if len(repoArr) == 2 {
			queryPara2.repo = repoArr[1]
		} else {
			queryPara2.month = monthArr[1]
		}

		// 获取结果
		ret1 = getResult(queryPara1)
		ret2 = getResult(queryPara2)

		// 汇总结果到ret里

		ret := map[string]string{
			"repo.name": ret1["repo.name"],
			"repo.url":  ret1["repo.url"],
		}

		for k,v1 := range(ret1) {
			if (k == "repo.name" || k == "repo.url"){
				continue
			}
			v2 := ret2[k]
			ret[k] = getDiffValue(v1,v2)
		}
		source = ret
	},
}

// 计算差值(左减右)，目前默认所有指标的值都是浮点型
func getDiffValue(a string,b string)string{
	intVal1,_ := strconv.ParseFloat(a, 10);
	intVal2,_ := strconv.ParseFloat(b, 10);
	return strconv.FormatFloat(intVal1 - intVal2, 'f', -1, 32)
}

func init() {
	compareCmd.Flags().StringArrayVarP(&repoArr, "repo", "r", nil, "Repository asked")
	compareCmd.MarkFlagRequired("repo")
	compareCmd.Flags().StringArrayVarP(&monthArr, "month", "M", nil, "Month of asked metric")
	compareCmd.MarkFlagRequired("month")
	compareCmd.Flags().StringVarP(&queryPara.metric, "metric", "m", "", "Metric asked")


	var showCompareCmd = *compareCmd
	var downloadCompareCmd = *compareCmd
	showCmd.AddCommand(&showCompareCmd)
	downloadCmd.AddCommand(&downloadCompareCmd)
}
