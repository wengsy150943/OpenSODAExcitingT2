/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"strings"
	"testing"
)




func cleanStateCompare(){
	repoArr = nil
	userArr = nil
	monthArr = nil
	compareCmd.ResetFlags()
	compareCmd.Flags().StringArrayVarP(&repoArr, "repo", "r", nil, "Repositories asked")
	compareCmd.Flags().StringArrayVarP(&userArr, "user", "u", nil, "Users asked")
	compareCmd.MarkFlagsMutuallyExclusive("repo", "user")

	compareCmd.Flags().StringArrayVarP(&monthArr, "month", "M", nil, "Month of asked metric")
	compareCmd.Flags().StringVarP(&queryPara.metric, "metric", "m", "", "Metric asked")

	var showCompareCmd = *compareCmd
	var downloadCompareCmd = *compareCmd

	showCmd.ResetCommands()
	downloadCmd.ResetCommands()

	showCmd.AddCommand(&showCompareCmd)
	downloadCmd.AddCommand(&downloadCompareCmd)
}

// illegal case
func TestComparePanic(t *testing.T) {
	

	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()


	var basic_case = []string{
		"show ",
		"download ",
	}

	var test_case = []string{
		"compare -r X-lab2017/open-digger -M 2023-04 -m openrank",
		"compare -r X-lab2017/open-digger -M 2023-04 -m openrank -m openrank1",
		"compare -r X-lab2017/open-digger -m openrank",
		"compare -M 2023-04 -m openrank",
		"compare -r X-lab2017/open-digger -M 2023-04 -M 2023-04 -M 2023-04 -m openrank",
		"compare -r X-lab2017/open-digger -r X-lab2017/open-digger -r X-lab2017/open-digger -M 2023-04 -m openrank",
		"compare -u X-lab2017 -u X-lab2017 -u X-lab2017 -M 2023-04 -m openrank",
		"compare -r X-lab2017/open-digger -u X-lab2017 -M 2023-04 -M 2023-04 -m openrank",
	}

	for _,bv := range basic_case{
		for _,v := range test_case {
			cleanStateCompare()
			compareCmd.SetArgs(strings.Split(bv + v," "))
			compareCmd.Execute()
		}	
	}

	
}

func TestCompare(t *testing.T) {
	var basic_case  = []string{"show compare -m openrank", "download compare -m openrank"}
	var test_case = []string{
		" -r X-lab2017/open-digger -M 2023-05 -M 2023-04 ",
		" -r X-lab2017/open-digger -r X-lab2017/open-digger -M 2023-04 ",
		" -u X-lab2017 -u X-lab2018",
	}
	for _,bv := range basic_case{
		for _,v := range test_case {
			cleanStateCompare()
			compareCmd.SetArgs(strings.Split(bv + v," "))
			compareCmd.Execute()
		}	
	}
	
	
}


