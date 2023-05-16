/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
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
		// 获取结果
		source = map[string]string{
			"repo":strings.Join(repoArr,""),
			"month":strings.Join(monthArr,""),
		}
		fmt.Print(source)
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
