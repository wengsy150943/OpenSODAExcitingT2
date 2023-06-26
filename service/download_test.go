package service

import (
	"fmt"
	"testing"
)

func TestSingleDownloadService(t *testing.T) {
	fmt.Println("TestDownloadAsHtml：")
	var downloadService DownloadService
	downloadService = &SingleDownloadService{}
	downloadService.SetData("www.github.com/aaaaa", "html_output")
	downloadService.Download()
}

//func TestDownloadAsPdf(t *testing.T) {
//	fmt.Println("TestDownloadAsPdf：")
//	var downloadService DownloadService
//	downloadService = &DownloadData{}
//	downloadService.SetData("www.github.com/aaaaa")
//	downloadService.DownloadAsPdf("output_pdf")
//}

func TestBatchDownloadService(t *testing.T) {
	fmt.Println("TestDownloadAsHtml：")
	var downloadService DownloadService
	downloadService = &BatchDownloadService{}
	downloadService.SetData("aaaa.txt", "csv_output")
	downloadService.Download()
}
