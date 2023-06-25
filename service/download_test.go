package service

import (
	"fmt"
	"testing"
)

func TestDownloadAsHtml(t *testing.T) {
	fmt.Println("TestDownloadAsHtml：")
	var downloadService DownloadService
	downloadService = &DownloadData{}
	downloadService.SetData("www.github.com/aaaaa")
	downloadService.DownloadAsHtml("output_html")
}

func TestDownloadAsPdf(t *testing.T) {
	fmt.Println("TestDownloadAsPdf：")
	var downloadService DownloadService
	downloadService = &DownloadData{}
	downloadService.SetData("www.github.com/aaaaa")
	downloadService.DownloadAsPdf("output_pdf")
}
