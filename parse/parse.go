/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"exciting-opendigger/service"
)

// 根据查询参数获取结果
func getResult(QueryPara Query) service.RepoInfo {
	// 正确性验证
	if QueryPara.month == "" && QueryPara.metric == "" {
		if isShow {
			panic("Lack of enough parameters: Either metric or month must be specified.")
		}

		return service.GetAllRepoInfo(QueryPara.repo)
	}
	// 特定指标
	if QueryPara.month == "" {
		return service.GetRepoInfoOfMetric(QueryPara.repo, QueryPara.metric)
	}

	if QueryPara.metric == "" { // 特定月份的整体报告
		return service.GetRepoInfoOfMonth(QueryPara.repo, QueryPara.month)
	}

	{ // 特定月份在特定指标上的数据
		if !isShow {
			panic("Lack of enough data: Too much parameters. One of metric or month is enough.")
		}

		return service.GetCertainRepoInfo(QueryPara.repo, QueryPara.metric, QueryPara.month)
	}
}
