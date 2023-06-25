package service

import (
	pdf "github.com/adrg/go-wkhtmltopdf"
	"log"
	"os"
	"text/template"
)

type DownloadService interface {
	SetData(api_ string) error
	DownloadAsHtml(target_ string) error
	DownloadAsPdf(target_ string) error
}

type DownloadData struct {
	Api       string
	Title     string
	Dates     []string
	Openranks []float64
	Recent    map[string]float64
}

func (d *DownloadData) SetData(api_ string) error {
	//这里需要调用lhg的数据获取接口,数据结构DownloadData也根据接口数据修改！！！！！！！！！！！！！
	d.Api = api_
	d.Title = "opendigger"
	dates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	d.Dates = append(d.Dates, dates...)
	openranks := []float64{4.5, 4.91, 5.59, 6.31, 9.96, 10.61, 6.28, 4.14, 4.44, 4.26, 6.46, 4.84, 3.93, 3.34, 3, 2.84, 2.89, 3.33, 4.71, 4.87, 6.06, 3.76, 4.14, 7.67, 9.17, 8.53, 9.96, 11.84, 14.65, 19.36, 19.9, 40.48, 22.05, 18.79}
	d.Openranks = append(d.Openranks, openranks...)

	d.Recent = make(map[string]float64)

	for i := len(dates) - 1; i >= len(dates)-5; i-- {
		if i < 0 {
			break
		}
		d.Recent[dates[i]] = openranks[i]
	}

	return nil
}

func (d *DownloadData) DownloadAsHtml(target_ string) error {

	tmpl, err := template.ParseFiles("../assets/template/template.html")
	if err != nil {
		log.Fatal(err)
	}

	// 创建输出文件
	file, err := os.Create(target_ + ".html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	page := d

	// 渲染模板并将结果写入文件
	err = tmpl.Execute(file, page)
	if err != nil {
		panic(err)
	}

	return nil
}

func (d *DownloadData) DownloadAsPdf(target_ string) error {
	d.DownloadAsHtml(target_)

	// 初始化PDF实例
	if err := pdf.Init(); err != nil {
		log.Fatal(err)
	}
	defer pdf.Destroy()

	// 从HTML创造 object
	inFile, err := os.Open(target_ + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	object3, err := pdf.NewObjectFromReader(inFile)
	if err != nil {
		log.Fatal(err)
	}
	object3.Zoom = 1.5
	object3.TOC.Title = "Table of Contents"

	// 创造converter
	converter, err := pdf.NewConverter()
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()

	// 添加object
	converter.Add(object3)

	// PDF设置
	converter.Title = "Sample document"
	converter.PaperSize = pdf.A4
	converter.Orientation = pdf.Landscape
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "10mm"
	converter.MarginRight = "10mm"

	// 转变为PDF
	outFile, err := os.Create(target_ + ".pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Run converter.
	if err := converter.Run(outFile); err != nil {
		log.Fatal(err)
	}
	return nil
}
