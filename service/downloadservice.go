package service

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
)

type DownloadService interface {
	SetData(source_ RepoInfo, target_ string) error
	Download() error
	//DownloadAsPdf(target_ string) error
}

type SingleDownloadService struct {
	Source    string
	Target    string
	Title     string
	Dates     []string
	OpenRanks []float64
	Recent    map[string]float64
}

func (d *SingleDownloadService) SetData(source_ RepoInfo, target_ string) error {
	d.Target = target_
	//这里需要调用lhg的数据获取接口,数据结构SingleDownloadService也根据接口数据修改！！！！！！！！！！！！！
	d.Source = source_.repoName
	d.Title = "opendigger"
	dates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	d.Dates = append(d.Dates, dates...)
	openRanks := []float64{4.5, 4.91, 5.59, 6.31, 9.96, 10.61, 6.28, 4.14, 4.44, 4.26, 6.46, 4.84, 3.93, 3.34, 3, 2.84, 2.89, 3.33, 4.71, 4.87, 6.06, 3.76, 4.14, 7.67, 9.17, 8.53, 9.96, 11.84, 14.65, 19.36, 19.9, 40.48, 22.05, 18.79}
	d.OpenRanks = append(d.OpenRanks, openRanks...)

	d.Recent = make(map[string]float64)

	for i := len(dates) - 1; i >= len(dates)-5; i-- {
		if i < 0 {
			break
		}
		d.Recent[dates[i]] = openRanks[i]
	}

	return nil
}

func (d *SingleDownloadService) Download() error {

	tmpl, err := template.ParseFiles("../assets/template/template.html")
	if err != nil {
		log.Fatal(err)
	}

	// 创建输出文件
	file, err := os.Create(d.Target + ".html")
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

//
//func (d *DownloadData) DownloadAsPdf(target_ string) error {
//	d.DownloadAsHtml(target_)
//
//	// 初始化PDF实例
//	if err := pdf.Init(); err != nil {
//		log.Fatal(err)
//	}
//	defer pdf.Destroy()
//
//	// 从HTML创造 object
//	inFile, err := os.Open(target_ + ".html")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer inFile.Close()
//
//	object3, err := pdf.NewObjectFromReader(inFile)
//	if err != nil {
//		log.Fatal(err)
//	}
//	object3.Zoom = 1.5
//	object3.TOC.Title = "Table of Contents"
//
//	// 创造converter
//	converter, err := pdf.NewConverter()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer converter.Destroy()
//
//	// 添加object
//	converter.Add(object3)
//
//	// PDF设置
//	converter.Title = "Sample document"
//	converter.PaperSize = pdf.A4
//	converter.Orientation = pdf.Landscape
//	converter.MarginTop = "1cm"
//	converter.MarginBottom = "1cm"
//	converter.MarginLeft = "10mm"
//	converter.MarginRight = "10mm"
//
//	// 转变为PDF
//	outFile, err := os.Create(target_ + ".pdf")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer outFile.Close()
//
//	// Run converter.
//	if err := converter.Run(outFile); err != nil {
//		log.Fatal(err)
//	}
//	return nil
//}

type BatchDownloadData struct {
	BatchData []float64
}

type BatchDownloadService struct {
	Source RepoInfo
	Target string
	Rows   int
	Cols   int
	Data   map[string]BatchDownloadData
	Dates  []string
}

func (d *BatchDownloadService) SetData(source_ RepoInfo, target_ string) error {
	d.Target = target_
	//这里需要调用lhg的批量数据获取接口,数据结构BatchDownloadService和BatchDownloadData也根据接口数据修改！！！！！！！！！！！！！
	d.Source = source_
	d.Data = make(map[string]BatchDownloadData)

	dates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}

	d.Dates = append(d.Dates, dates...)

	openRanks := []float64{4.5, 4.91, 5.59, 6.31, 9.96, 10.61, 6.28, 4.14, 4.44, 4.26, 6.46, 4.84, 3.93, 3.34, 3, 2.84, 2.89, 3.33, 4.71, 4.87, 6.06, 3.76, 4.14, 7.67, 9.17, 8.53, 9.96, 11.84, 14.65, 19.36, 19.9, 40.48, 22.05, 18.79}

	data1 := &BatchDownloadData{}
	data1.BatchData = append(data1.BatchData, openRanks...)

	data2 := &BatchDownloadData{}
	data2.BatchData = append(data2.BatchData, openRanks...)

	d.Data["opendigger1"] = *data1
	d.Data["opendigger2"] = *data2

	d.Rows = 2
	d.Cols = len(dates)
	return nil
}

func (d *BatchDownloadService) Download() error {

	data := make([][]string, 0)

	firstRow := make([]string, 0)
	firstRow = append(firstRow, "仓库名")
	firstRow = append(firstRow, d.Dates...)

	data = append(data, firstRow)

	for key, value := range d.Data {
		tempRow := make([]string, d.Cols+1)
		tempRow[0] = key
		for j := 1; j <= d.Cols; j++ {
			tempRow[j] = strconv.FormatFloat(value.BatchData[j-1], 'f', 3, 64)
		}
		data = append(data, tempRow)
	}

	file, err := os.Create(d.Target + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}
	return nil
}
