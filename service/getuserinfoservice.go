package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserInfo struct {
	Username string
	Data     map[string](map[string]interface{})
}

func GetUrlContent(url, username, metric string) UserInfo {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 解析数据
	body, _ := ioutil.ReadAll(response.Body)

	var temp map[string]interface{}
	data_list := map[string]interface{}{}
	json.Unmarshal([]byte(body), &temp)

	//用户信息返回的结构类型为：map[string][][]sting{}
	// 获取日期并排序, 需要针对特殊情况做处理
	if metric == "developernetwork" || metric == "reponetwork" {
		nodesvalue, ok := temp["nodes"].([][]string)
		if !ok {
			panic("cannot convert temp to [][]string")
		}
		edgesvalue, _ := temp["edges"].([][]string)
		data_list["nodes"] = nodesvalue
		data_list["edges"] = edgesvalue
	} else {
		data_list = temp
	}
	// 将数据赋值给RepoInfo，如果数据是9种特殊指标，解析为specialData；并赋给data
	// 获取特殊指标对应的解析函数
	var data map[string](map[string]interface{})
	data = make(map[string](map[string]interface{}))
	data[metric] = data_list
	ret := UserInfo{
		Username: username,
		Data:     data,
	}

	return ret
}
func GetCertainUser(username string) {
	//TODO 获得四个指标的url，开四个携程分别去请求url，结果写入UserInfo
	BaseURL := "https://oss.x-lab.info/open_digger/github/"
	openrankURL := BaseURL + username + "/" + "openrank.json"
	developernetworkURL := BaseURL + username + "/" + "developer_network.json"
	GetUrlCotent(openrankURL)
}
