package service

import (
	"encoding/json"
	"errors"
	"exciting-opendigger/utils"
	"gorm.io/gorm"
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
	//判断是否已经创建了缓存表
	exists := utils.TableExist("cached_repo_infos")
	if !exists {
		utils.CreateTable()
	}
	cachedrepoinfo := utils.CachedRepoInfo{}
	repoName := strings.Split(repo, "/")[1]
	//先去缓存中查询该repo的信息是否被缓存
	err := utils.Readquerysinglemetric(&cachedrepoinfo, repoName, metric)
	//若缓存在sqlite中，则将缓存的值返回
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		ret := RepoInfo{
			repoName: cachedrepoinfo.Reponame,
			repoUrl:  cachedrepoinfo.Repourl,
			month:    "",
			data:     cachedrepoinfo.Data,
			dates:    cachedrepoinfo.Dates,
		}
		return ret
	}

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	repoURL := "https://github.com/" + repo

	var temp map[string]interface{}
	json.Unmarshal([]byte(body), &temp)

	// 获取日期并排序
	dates := make([]string, len(temp))
	cnt := 0
	for i := range temp {
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
	//查询结果插入缓存
	utils.Insertsinglequery(repoName, repoURL, metric, "", dates, data)
	return ret
}

// TODO 支持从缓存中查找指定 repo metric moth条目
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

// TODO
func GetRepoInfoOfMonth(repo, month string) RepoInfo {
	return RepoInfo{
		repoName: "",
		repoUrl:  "",
		month:    "",
		data:     nil,
	}
}
