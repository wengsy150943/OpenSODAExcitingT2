/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"exciting-opendigger/service"
	"strings"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show data from OpenDigger",
	Long:  `show data from api and print in screen`,
	PersistentPreRun: func (cmd *cobra.Command, args []string)  {
		isShow = true
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 获取结果
		if queryPara.user != "" {
			userInfo = service.GetCertainUser(queryPara.user)
		} else{
			repoInfo = getResult(queryPara)
		}
		
		
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		

		// 检查是否有compare,这里只有一个subcommand
		if (strings.Contains(cmd.CommandPath(), "compare") ){

			if repoInfo.RepoName != "" {
				service.PrintRepoInfo(repoInfo)
				service.PrintRepoInfo(repoInfoCompare)
			} else{
				service.PrintUserInfo(userInfo)
				service.PrintUserInfo(userInfoCompare)
			}
			
		} else{
			if repoInfo.RepoName != "" {
				service.PrintRepoInfo(repoInfo)
			} else{
				service.PrintUserInfo(userInfo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
