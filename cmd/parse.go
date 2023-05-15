/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"exciting-opendigger/service"
	"fmt"
)



// 根据查询参数获取结果
func getResult(QueryPara Query)map[string]string {
	// 正确性验证
	if QueryPara.month == "" && QueryPara.metric == "" {
		fmt.Errorf("Lack of enough parameters: either month or metric is required")
		return nil
	}

	// 特定指标
	if QueryPara.month == "" {
		ret, _, _ := service.GetCertainRepo(QueryPara.repo, QueryPara.metric)
		return ret
	}

	// 特定月份的整体报告
	if QueryPara.metric == "" {
		
	}

	// 特定月份在特定指标上的数据
	return service.GetCertainMonth(QueryPara.repo, QueryPara.metric, QueryPara.month)
}
