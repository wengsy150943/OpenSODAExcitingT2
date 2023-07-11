package service

import (
	"encoding/json"
	"errors"
	"exciting-opendigger/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"sort"
	"sync"
	"time"
)

type UserInfo struct {
	Username string
	Dates    []string
	Data     map[string](map[string]interface{})
}

func GetUserUrlContent(url, username string) map[string]interface{} {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 解析数据
	body, _ := ioutil.ReadAll(response.Body)
	var temp map[string]interface{}
	json.Unmarshal([]byte(body), &temp)
	return temp
}
func GetContentParal(res map[string](map[string]interface{}), Urls []string, username string, UserMetric []string) UserInfo {
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go appendlist(&wg, res, GetUserUrlContent(Urls[i], username), UserMetric[i])
	}
	wg.Wait()
	//println(res["reponetwork"]["nodes"].([]interface{})[0].([]interface{})[0].(string))
	//map赋值前需要先初始化(包括内部的map[string]interface{}
	data := make(map[string](map[string]interface{}))
	data["developernetwork"] = make(map[string]interface{})
	data["reponetwork"] = make(map[string]interface{})

	//包括内部的map[string]interface{}
	data_list := make([]map[string]interface{}, 4)
	for i := 0; i < 4; i++ {
		data_list[i] = make(map[string]interface{})
	}
	for k, _ := range res {
		//=使用浅赋值，data[k]指向data_list的地址，需要重新深拷贝
		if k == "openrank" {
			data_list[0] = res[k]
			data[k] = data_list[0]
			//data_list = res[k]
			//data[k] = data_list
		} else if k == "activity" {
			data_list[1] = res[k]
			data[k] = data_list[1]
		} else if k == "reponetwork" {
			//println(k)
			data_list[2]["nodes"] = res[k]["nodes"]
			data_list[2]["edges"] = res[k]["edges"]
			data[k] = data_list[2]
		} else {
			data_list[3]["nodes"] = res[k]["nodes"]
			data_list[3]["edges"] = res[k]["edges"]
			data[k] = data_list[3]
		}

	}
	dates := make([]string, len(data_list))
	cnt := 0
	for i := range data_list[0] {
		dates[cnt] = i
		cnt++
	}
	sort.Slice(dates, func(i, j int) bool { return dates[i] < dates[j] })
	ret := UserInfo{}
	ret.Dates = dates
	ret.Data = data

	return ret
}
func appendlist(wg *sync.WaitGroup, l map[string](map[string]interface{}), t map[string]interface{}, metric string) {
	l[metric] = t
	wg.Done()
}
func GetCertainUser(username string) UserInfo {
	//TODO 获得四个指标的url，开四个携程分别去请求url，结果写入UserInfo
	BaseURL := "https://oss.x-lab.info/open_digger/github/"
	openrankURL := BaseURL + username + "/" + "openrank.json"
	activity := BaseURL + username + "/" + "activity.json"
	developernetworkURL := BaseURL + username + "/" + "developer_network.json"
	reponetwork := BaseURL + username + "/" + "repo_network.json"
	ret := UserInfo{}
	ret.Username = username
	Urls := []string{openrankURL, activity, developernetworkURL, reponetwork}
	res := map[string](map[string]interface{}){}
	UserMetric := []string{"openrank", "activity", "developernetwork", "reponetwork"}

	exists := utils.TableExist("cached_user_infos")
	if !exists {
		utils.CreateTable(utils.CachedUserInfo{})
	}
	cacheduserinfo := utils.CachedUserInfo{}
	err := utils.ReadSingleUserInfo(&cacheduserinfo, username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		currentTime := time.Now()
		updateTime := cacheduserinfo.UpdatedAt
		duration := currentTime.Sub(updateTime)
		//更新时间超过24小时则重新获取并更新缓存
		if duration > 24*time.Hour {
			temp := GetContentParal(res, Urls, username, UserMetric)
			err := utils.UpdateUserInfoSingleRow(username, temp.Data, temp.Dates)
			if err != nil {
				panic("update" + username + " faild")
			}
		}
		ret := UserInfo{
			Username: cacheduserinfo.Username,
			Data:     cacheduserinfo.Data,
		}
		return ret
	}
	ret = GetContentParal(res, Urls, username, UserMetric)
	utils.InsertUserInfo(username, ret.Data, ret.Dates)
	return ret
}
