package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type UserInfo struct {
	Username string
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
func appendlist(wg *sync.WaitGroup, m *sync.Mutex, l map[string](map[string]interface{}), t map[string]interface{}, metric string) {
	m.Lock()
	l[metric] = t
	m.Unlock()
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

	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go appendlist(&wg, &m, res, GetUserUrlContent(Urls[i], username), UserMetric[i])
	}
	wg.Wait()
	//println(res["developernetwork"]["nodes"].([]interface{})[0].([]interface{})[0].(string))
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
	ret.Data = data

	return ret
}
