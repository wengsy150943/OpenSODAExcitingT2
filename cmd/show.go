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
	Short: "show data from api",
	Long:  `show data from api and print in screen`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
		// 获取结果
		source := getResult(queryPara)

		// TODO 输出结果，格式尚未整理
		fmt.Print(source)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
