package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//type RepoInfo struct {
//	name          string `json:"repo.name"`
//	URL           string `json:"repo.url"`
//	month         string
//	openRank      string `json:"repo.index.xlab.openrank"`
//	activity      string `json:"repo.index.xlab.activity"`
//	datesAndTimes string `json:"repo.metric.chaoss.active dates and times"`
//}

/*
*
Get_On_certain_repo
*/
func GetCertainRepo(repo string, metric string) (map[string]string, []byte, string) {

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

	return repoInfo, body, metric
}

func GetCertainMonth(repo string, metric string, month string) map[string]string {

	hashData, body, _ := GetCertainRepo(repo, metric)
	var v2 interface{}

	json.Unmarshal(body, &v2)
	data := v2.(map[string]interface{})
	repoInfo := map[string]string{}
	for k, v := range data {
		if k == month {
			repoInfo = map[string]string{
				"repo.name": hashData["repo.name"],
				"repo.url":  hashData["repo.url"],
				"month":     month,
				metric:      strconv.FormatFloat(v.(float64), 'f', 2, 32),
			}
		}
	}
	//bytes, _ := json.Marshal(repoInfo)
	return repoInfo
}
