package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//	type RepoInfo struct {
//		name          string `json:"repo.name"`
//		URL           string `json:"repo.url"`
//		month         string
//		openRank      string `json:"repo.index.xlab.openrank"`
//		activity      string `json:"repo.index.xlab.activity"`
//		datesAndTimes string `json:"repo.metric.chaoss.active dates and times"`
//	}
type RepoInfoservice interface {
	Getrepoinfo(repo string, metric string, month string)
}

type RepoInfo struct {
	metric   string
	reponame string
	repourl  string
	data     []byte
}
type RepoInfoMonth struct {
	metric   string
	reponame string
	repourl  string
	data     string
	month    string
}

func (r *RepoInfo) Getrepoinfo(repo, metric, month string) {
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

	r.metric = metric
	r.reponame = repoName
	r.repourl = repoURL
	r.data = body
}

func (r *RepoInfoMonth) Getrepoinfo(repo, metric, month string) {
	a := RepoInfo{}
	a.Getrepoinfo(repo, metric, "")

	r.repourl = a.repourl
	r.reponame = a.reponame
	r.metric = a.metric

	body := a.data
	var temp interface{}
	json.Unmarshal([]byte(body), &temp)
	d := temp.(map[string]interface{})
	for k, v := range d {
		if k == month {
			println(month)
			r.data = strconv.FormatFloat(v.(float64), 'f', 2, 32)
			r.month = month
		}
	}
}
