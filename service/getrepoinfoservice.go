package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type RepoInfo struct {
	name          string `json:"repo.name"`
	URL           string `json:"repo.url"`
	month         string
	openRank      string `json:"repo.index.xlab.openrank"`
	activity      string `json:"repo.index.xlab.activity"`
	datesAndTimes string `json:"repo.metric.chaoss.active dates and times"`
}

/*
*
Get_On_certain_repo
*/
func GetCertainRepo(repo string, metric string) ([]byte, []byte) {

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
	repoInfo := map[string]string{
		"repo.name": repoName,
		"repo.url":  repoURL,
		metric:      string(body),
	}
	bytes, _ := json.Marshal(repoInfo)

	return bytes, body
}

func GetCertainMonth(repo string, metric string, month string) []byte {

	jsonData, body := GetCertainRepo(repo, metric)
	var v1 interface{}
	var v2 interface{}

	json.Unmarshal(jsonData, &v1)
	json.Unmarshal(body, &v2)
	data1 := v1.(map[string]interface{})
	data2 := v2.(map[string]interface{})
	repoInfo := map[string]string{}
	for k, v := range data2 {
		if k == month {
			repoInfo = map[string]string{
				"repo.name": data1["repo.name"].(string),
				"repo.url":  data1["repo.url"].(string),
				"month":     month,
				metric:      strconv.FormatFloat(v.(float64), 'f', 2, 32),
			}
		}
	}
	bytes, _ := json.Marshal(repoInfo)
	return bytes
}
