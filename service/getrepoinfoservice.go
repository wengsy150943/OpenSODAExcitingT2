package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"sync"
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

func GetRepoInfoOfMonth(repo, month string) (repoinfo RepoInfo) {
	MetricPerThread := MetricNum / GoroutinueNum
	var repoinfos [MetricNum]RepoInfo
	var begin, end int
	id := 0
	var wg sync.WaitGroup

	for id < GoroutinueNum {
		wg.Add(1)
		// 划定每个协程处理的范围
		begin = id * MetricPerThread
		if id == GoroutinueNum-1 {
			end = MetricNum
		} else {
			end = (id + 1) * MetricPerThread
		}

		go func(begin int, end int) {
			for i := begin; i < end; i++ {
				repoinfos[i] = GetCertainRepoInfo(repo, Metrics[i], month)
			}
			wg.Done()
		}(begin, end)
		id++
	}
	wg.Wait()

	dateMap := map[string]bool{}
	repoinfo.repoName = repo
	repoinfo.repoUrl = repoinfos[0].repoUrl
	repoinfo.month = month
	repoinfo.data = make(map[string](map[string]interface{}))

	for i := 0; i < MetricNum; i++ {
		for _, date := range repoinfos[i].dates {
			dateMap[date] = true
		}
		repoinfo.data[Metrics[i]] = repoinfos[i].data[Metrics[i]]
	}

	dates := make([]string, len(dateMap))
	cnt := 0
	for k, _ := range dateMap {
		dates[cnt] = k
		cnt++
	}

	sort.Slice(dates, func(i, j int) bool { return dates[i] < dates[j] })
	repoinfo.dates = dates
	return
}
