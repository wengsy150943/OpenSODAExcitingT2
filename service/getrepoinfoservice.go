package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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
	month    string
	data     []byte
}

func (r *RepoInfo) Getrepoinfo(repo, metric string) {
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
	a.Getrepoinfo(repo, metric)

	r.repourl = a.repourl
	r.reponame = a.reponame
	r.metric = a.metric
	r.data = []byte{}

	body := a.data
	var temp interface{}
	json.Unmarshal([]byte(body), &temp)
	d := temp.(map[string]interface{})
	value, ok := d[month]
	if ok {
		r.data = append(r.data, []byte(fmt.Sprintf("%v", value))...)
	} else {
		for k, v := range d {
			switch v.(type) {
			case map[string]interface{}:
				innermap := v.(map[string]interface{})
				v2, ok := innermap[month]
				if ok {
					r.data = append(r.data, []byte(fmt.Sprintf("%v:%v ", k, v2))...)
				}
			}
		}
	}
}

func GetCertainMonth(repo string, month string) map[string]string{
	repoInfo := map[string]string{}
	return repoInfo
}