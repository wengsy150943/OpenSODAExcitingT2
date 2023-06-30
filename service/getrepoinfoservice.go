package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type RepoInfo struct {
	repoName string
	repoUrl  string
	month    string
	data     map[string](map[string]float32)
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

	var temp map[string]float32
	json.Unmarshal([]byte(body), &temp)

	

	data := make(map[string](map[string]float32))
	data[metric] = temp

	ret := RepoInfo{
		repoName: repoName,
		repoUrl:  repoURL,
		month:    "",
		data:     data,
	}

	return ret
}

func GetCertainRepoInfo(repo, metric, month string) RepoInfo {
	repoInfo := GetRepoInfoOfMetric(repo, metric)
	repoInfo.month = month

	data := make(map[string](map[string]float32))

	for k, v := range repoInfo.data {
		data[k] = map[string]float32{month: v[month]}
	}

	repoInfo.data = data

	return repoInfo
}

func GetRepoInfoOfMonth(repo, month string) RepoInfo{
	return RepoInfo{
		repoName: "",
		repoUrl:  "",
		month:    "",
		data:     nil,
	}
}
