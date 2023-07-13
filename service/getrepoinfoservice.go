package service

import (
	"encoding/json"
	"errors"
	"exciting-opendigger/utils"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type RepoInfo struct {
	RepoName    string
	RepoUrl     string
	Month       string
	Dates       []string
	Data        map[string](map[string]interface{})
	SpecialData utils.SpecialDataStructure
}

// 把特殊metric转存一份到specialData里
func initSpecialDataStructure(data map[string]map[string]interface{}) utils.SpecialDataStructure {
	var specialData utils.SpecialDataStructure
	for k, v := range data {
		parseFunction, ok := utils.Parse[k]
		if ok {
			specialData = parseFunction(v, specialData)
		}
	}
	return specialData
}

func GetUrlContent(url string, repo string, metric string) RepoInfo {
	repoName := strings.Split(repo, "/")[1]
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 解析数据
	body, _ := ioutil.ReadAll(response.Body)
	repoURL := "https://github.com/" + repo

	var temp map[string]interface{}
	data_list := map[string]interface{}{}
	json.Unmarshal([]byte(body), &temp)

	// 获取日期并排序, 需要针对特殊情况做处理
	if Special_Metric[metric] {
		data_list = temp["avg"].(map[string]interface{})
	} else {
		data_list = temp
	}

	dates := make([]string, len(data_list))
	cnt := 0
	for i := range data_list {
		dates[cnt] = i
		cnt++
	}
	sort.Slice(dates, func(i, j int) bool { return dates[i] < dates[j] })

	// 将数据赋值给RepoInfo，如果数据是9种特殊指标，解析为specialData；并赋给data
	// 获取特殊指标对应的解析函数
	var data map[string](map[string]interface{})
	data = make(map[string](map[string]interface{}))
	data[metric] = temp

	ret := RepoInfo{
		RepoName:    repoName,
		RepoUrl:     repoURL,
		Month:       "",
		Data:        data,
		Dates:       dates,
		SpecialData: initSpecialDataStructure(data),
	}

	fmt.Println(repoName)

	return ret
}

func GetRepoInfoOfMetric(repo, metric string) RepoInfo {
	BaseURL := "https://oss.x-lab.info/open_digger/github/"
	url := BaseURL + repo + "/" + strings.ToLower(metric) + ".json"
	//判断是否已经创建了缓存表
	exists := utils.TableExist("cached_repo_infos")
	if !exists {
		utils.CreateTable(utils.CachedRepoInfo{})
	}
	cachedrepoinfo := utils.CachedRepoInfo{}
	repoName := strings.Split(repo, "/")[1]
	//先去缓存中查询该repo的信息是否被缓存
	err := utils.ReadQuerySingleMetric(&cachedrepoinfo, repoName, metric)
	//若缓存在sqlite中，则将缓存的值返回
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//判断缓存是否最新
		//将缓存的修改时间与当前时间做对比,
		currentTime := time.Now()
		updateTime := cachedrepoinfo.UpdatedAt
		duration := currentTime.Sub(updateTime)
		//更新时间超过24小时则重新获取并更新缓存
		if duration > 24*time.Hour {
			temp := GetUrlContent(url, repo, metric)
			err := utils.UpdateSingleRow(repoName, metric, temp.Dates, temp.Data)
			if err != nil {
				panic("update" + repoName + " " + metric + " faild")
			}
		}
		ret := RepoInfo{
			RepoName:    cachedrepoinfo.Reponame,
			RepoUrl:     cachedrepoinfo.Repourl,
			Month:       "",
			Data:        cachedrepoinfo.Data,
			Dates:       cachedrepoinfo.Dates,
			SpecialData: initSpecialDataStructure(cachedrepoinfo.Data),
		}
		return ret
	}

	ret := GetUrlContent(url, repo, metric)
	//查询结果插入缓存
	utils.InsertSingleQuery(repoName, ret.RepoUrl, metric, "", ret.Dates, ret.Data)

	fmt.Println(cachedrepoinfo.Reponame)
	fmt.Println(repoName)

	return ret
}

func GetCertainRepoInfo(repo, metric, month string) RepoInfo {
	repoInfo := GetRepoInfoOfMetric(repo, metric)
	repoInfo.Month = month

	data := make(map[string](map[string]interface{}))

	// 处理特殊指标
	_, ok := utils.Parse[metric]
	if ok {
		repoInfo.SpecialData.SelectMonth(month)
	}

	// 因为仍然保留一份数据在data里，这部分也要处理，用于show的输出
	if Special_Metric[metric] {
		for k, v := range repoInfo.Data {
			dataMap := make(map[string]interface{})
			for _, val := range Special_Value {
				dataMap[val] = v[val].(map[string]interface{})[month]
			}
			data[k] = map[string]interface{}{month: dataMap}
		}
	} else {
		for k, v := range repoInfo.Data {
			data[k] = map[string]interface{}{month: v[month]}
		}
	}

	repoInfo.Data = data
	repoInfo.Dates = []string{month}

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
	repoinfo.RepoName = repo
	repoinfo.RepoUrl = repoinfos[0].RepoUrl
	repoinfo.Month = month
	repoinfo.Data = make(map[string](map[string]interface{}))

	for i := 0; i < MetricNum; i++ {
		for _, date := range repoinfos[i].Dates {
			dateMap[date] = true
		}
		repoinfo.Data[Metrics[i]] = repoinfos[i].Data[Metrics[i]]
		repoinfo.SpecialData.MergeSpecialData(repoinfos[i].SpecialData)
	}

	dates := make([]string, len(dateMap))
	cnt := 0
	for k, _ := range dateMap {
		dates[cnt] = k
		cnt++
	}

	sort.Slice(dates, func(i, j int) bool { return dates[i] < dates[j] })
	repoinfo.Dates = dates
	return repoinfo
}

func GetAllRepoInfo(repo string) (repoinfo RepoInfo) {
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
				repoinfos[i] = GetRepoInfoOfMetric(repo, Metrics[i])
			}
			wg.Done()
		}(begin, end)
		id++
	}
	wg.Wait()

	dateMap := map[string]bool{}
	repoinfo.RepoName = repo
	repoinfo.RepoUrl = repoinfos[0].RepoUrl
	repoinfo.Data = make(map[string](map[string]interface{}))

	for i := 0; i < MetricNum; i++ {
		for _, date := range repoinfos[i].Dates {
			dateMap[date] = true
		}
		repoinfo.Data[Metrics[i]] = repoinfos[i].Data[Metrics[i]]
		repoinfo.SpecialData.MergeSpecialData(repoinfos[i].SpecialData)
	}

	dates := make([]string, len(dateMap))
	cnt := 0
	for k, _ := range dateMap {
		dates[cnt] = k
		cnt++
	}

	sort.Slice(dates, func(i, j int) bool { return dates[i] < dates[j] })
	repoinfo.Dates = dates
	return repoinfo
}
