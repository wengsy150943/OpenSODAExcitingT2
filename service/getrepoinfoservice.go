package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type RepoInfo struct {
	repoName string
	repoUrl  string
	month    string
	dates    []string
	data     map[string](map[string]interface{})
}

func GetRepoInfoOfMetric(repo, metric string) RepoInfo {
	BaseURL := "https://oss.x-lab.info/open_digger/github/"
	url := BaseURL + repo + "/" + strings.ToLower(metric) + ".json"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	repoName := strings.Split(repo, "/")[1]
	repoURL := "https://github.com/" + repo

	var temp map[string]interface{}
	data_list := &map[string]interface{}{}
	json.Unmarshal([]byte(body), &temp)

	// 获取日期并排序, 需要针对特殊情况做处理
	cnt := 0
	if Special_Metric[metric] {
		*data_list = temp["avg"].(map[string]interface{})
	} else {
		data_list = &temp
	}

	dates := make([]string, len(*data_list))
	for i := range *data_list {
		dates[cnt] = i
		cnt++
	}

	sort.Slice(dates, func(i, j int) bool { return dates[i] < dates[j] })

	data := make(map[string](map[string]interface{}))
	data[metric] = temp

	ret := RepoInfo{
		repoName: repoName,
		repoUrl:  repoURL,
		month:    "",
		data:     data,
		dates:    dates,
	}

	return ret
}

func GetCertainRepoInfo(repo, metric, month string) RepoInfo {
	repoInfo := GetRepoInfoOfMetric(repo, metric)
	repoInfo.month = month

	data := make(map[string](map[string]interface{}))

	for k, v := range repoInfo.data {
		data[k] = map[string]interface{}{month: v[month]}
	}

	repoInfo.data = data
	repoInfo.dates = []string{month}

	return repoInfo
}

func GetRepoInfoOfMonth(repo, month string) RepoInfo {
	return RepoInfo{
		repoName: "",
		repoUrl:  "",
		month:    "",
		data:     nil,
	}
}
