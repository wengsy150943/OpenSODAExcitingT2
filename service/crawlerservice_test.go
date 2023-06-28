package service

import (
	"fmt"
	"testing"
)

func TestCrawlerService(t *testing.T) {
	var crawlerService CrawlerService
	crawlerService = &CrawlTrendingService{}
	//options用于根据前端接受的参数设置爬虫选项，分别是日期、
	opts := make([]option, 0)
	opts = append(opts, WithDaily())
	opts = append(opts, WithProgramLanguage("java"))
	opts = append(opts, WithSpokenLanguage("chinese"))
	crawlerService.loadOptions(opts...)
	result, err := crawlerService.Crawl()

	if err != nil {
		fmt.Println(err)
	}
	for _, info := range result {
		fmt.Println(info.Link)
	}
}
