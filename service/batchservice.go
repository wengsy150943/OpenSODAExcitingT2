package service

import (
	"log"
	"sync"
)

func DownLoadRepoList(repoList []string, outputfile string) {
	repoInfoList := [MetricNum][]RepoInfo{}
	downloadList := make([]BatchDownloadService, GoroutinueNum)
	id := 0
	MetricPerThread := MetricNum / GoroutinueNum
	var begin, end int
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
		go func(begin int, end int, id int) {
			for i := begin; i < end; i++ {
				for _, repo := range repoList {
					repoInfoList[i] = append(repoInfoList[i], GetRepoInfoOfMetric(repo, Metrics[i]))
				}
				if err := downloadList[id].SetData(repoInfoList[i], Metrics[i], outputfile); err != nil {
					log.Fatal(err)
				}
				if err := downloadList[id].Download(); err != nil {
					log.Fatal(err)
				}
			}
			wg.Done()
		}(begin, end, id)
		id++
	}
	wg.Wait()

	return
}
