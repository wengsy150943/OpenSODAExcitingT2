/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show data from OpenDigger",
	Long:  `show data from api and print in screen`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取结果
		source = getResult(queryPara)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// TODO 输出结果，格式尚未整理
		if (source != nil){
			fmt.Print(source)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
