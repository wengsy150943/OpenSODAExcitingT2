package service

import (
	"github.com/signintech/gopdf"
	"log"
)

type DownloadService interface {
	SetData(api_ string) error
	Download(target_ string) error
}

type DownloadAsPdf struct {
	Api  string
	Data string
}

func (d *DownloadAsPdf) SetData(api_ string) error {
	d.Api = api_
	d.Data = "here is final answer from api" + d.Api + ":data data data"
	return nil
}

func (d *DownloadAsPdf) Download(target_ string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("微软雅黑", "assets/ttf/微软雅黑.ttf")
	if err != nil {
		log.Print(err.Error())
		return nil
	}

	err = pdf.SetFont("微软雅黑", "", 14)
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	pdf.Cell(nil, d.Data)
	pdf.WritePdf(target_)
	return nil
}
