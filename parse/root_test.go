/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"fmt"
	"strings"
	"testing"
)



func cleanStateRoot() {

	rootCmd.ResetFlags()
	rootCmd.ResetCommands()

	queryPara.metric = ""
	queryPara.month = ""
	queryPara.repo = ""
	queryPara.user = ""
	
	rootCmd.PersistentFlags().StringVarP(&queryPara.repo, "repo", "r", "", "Repository asked")
	rootCmd.PersistentFlags().StringVarP(&queryPara.user, "user", "u", "", "User asked")

	// rootCmd.MarkFlagsMutuallyExclusive("repo", "user")

	rootCmd.PersistentFlags().StringVarP(&queryPara.month, "month", "M", "", "Month of asked metric")
	rootCmd.PersistentFlags().StringVarP(&queryPara.metric, "metric", "m", "", "Metric asked")

	// 因为user不支持任何参数，在传参的时候预先卡死
	// rootCmd.MarkFlagsMutuallyExclusive("user", "month")
	// rootCmd.MarkFlagsMutuallyExclusive("user", "metric")

	// download

	position = ""
	downloadCmd.ResetFlags()
	downloadCmd.Flags().StringVarP(&position, "position", "p", "", "Download place where data to write to")
	err := downloadCmd.MarkFlagRequired("position")
	if err != nil {
		panic("Fail to MarkFlagRequired")
	}
	downloadCmd.Flags().BoolVarP(&user, "user", "u", false, "download user's data")

	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(logCmd)
}


func TestRootShow(t *testing.T) {

	var basic_case  = []string{"show "}
	var test_case = []string{
		" -r X-lab2017/open-digger -M 2023-05  -m openrank",
		" -r X-lab2017/open-digger -m openrank",
		" -r X-lab2017/open-digger -M 2023-05 ",
		" -u X-lab2017",
	}
	for _,bv := range basic_case{
		for _,v := range test_case {
			cleanStateRoot()
			isShow = true
			fmt.Println(bv + v)
			rootCmd.SetArgs(strings.Split(bv + v," "))
			rootCmd.Execute()
		}	
	}
}

func TestRootShowCompare(t *testing.T) {

	var basic_case  = []string{"show compare "}
	var test_case = []string{
		" -r X-lab2017/open-digger -M 2023-05 -M 2023-05  -m openrank",
		" -r X-lab2017/open-digger -r X-lab2017/open-digger -M 2023-05 ",
		" -u X-lab2017 -u X-lab2017",
	}
	for _,bv := range basic_case{
		for _,v := range test_case {
			cleanStateRoot()
			cleanStateCompare()
			fmt.Println(bv + v)
			rootCmd.SetArgs(strings.Split(bv + v," "))
			rootCmd.Execute()
		}	
	}
}

func TestRootOther(t *testing.T) {
	rootCmd.SetArgs(strings.Split("version"," "))
	rootCmd.Execute()

	rootCmd.SetArgs(strings.Split("log", " "))
	rootCmd.Execute()
}



// func TestRootDownload(t *testing.T) {

// 	var basic_case  = []string{" download -p a "}
// 	var test_case = []string{
// 		" -r X-lab2017/open-digger -M 2023-05  -m openrank",
// 		" -r X-lab2017/open-digger -M 2023-05 ",
// 		" -u X-lab2017",
// 	}
// 	for _,bv := range basic_case{
// 		for _,v := range test_case {
// 			cleanStateRoot()
// 			fmt.Println(bv + v)
// 			rootCmd.SetArgs(strings.Split(bv + v," "))
// 			rootCmd.Execute()
// 		}	
// 	}
// }
