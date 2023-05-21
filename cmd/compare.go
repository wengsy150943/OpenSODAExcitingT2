/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
		fmt.Println("compare called")
		if len(repoArr) + len(monthArr) != 3 {
			fmt.Errorf("argc number mismatch")
		}

		// 获取结果
		var ret1 map[string]string
		var ret2 map[string]string
		if len(repoArr) == 2 {
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
			ret1 = getResult(queryPara1)
			ret2 = getResult(queryPara2)
		} else {
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
			ret1 = getResult(queryPara1)
			ret2 = getResult(queryPara2)
		}
		fmt.Print(ret1)
		fmt.Print(ret2)

	},
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
