package service

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"strconv"
)

type SingleDownloadService struct {
	Source string
	Target string
	Title  string
	Dates  []string
	Data   map[string]([]float32)
}

func parseFloatValue(v interface{}) float32 {
	switch v.(type) {
		case float32: return v.(float32);
		case float64: return float32(v.(float64));
		case int: return float32(v.(int));
	}
	return v.(float32);
}

func (d *SingleDownloadService) SetData(source_ RepoInfo, target_ string) error {
	d.Target = target_
	d.Source = source_.repoUrl
	d.Title = source_.repoName

	d.Dates = source_.dates
	d.Data = make(map[string]([]float32))
	for k, v1 := range source_.data {
		tempList := make([]float32, 0)
		for _, v2 := range d.Dates {
			temp, ok := v1[v2]
			if ok {
				tempList = append(tempList, parseFloatValue(temp))
			} else {
				continue
			}
		}
		d.Data[k] = tempList
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

type BatchDownloadService struct {
	Metric string
	Target string
	Rows   int
	Cols   int
	Data   map[string]([]float32) //这里的key为仓库名
	Dates  []string
}

func (d *BatchDownloadService) SetData(sources_ []RepoInfo, metric_ string, target_ string) error {
	d.Target = target_
	d.Metric = metric_
	maxLength := 0
	for _, repo := range sources_ {
		if len(repo.dates) > maxLength {
			d.Dates = repo.dates
			maxLength = len(d.Dates)
		}
	}

	d.Data = make(map[string]([]float32))

	for _, repo := range sources_ {
		name := repo.repoName
		tempList := make([]float32, 0)
		for _, v := range d.Dates {
			temp, ok := repo.data[metric_][v]
			if ok {
				tempList = append(tempList, parseFloatValue(temp))
			} else {
				tempList = append(tempList, 0)
			}
		}
		d.Data[name] = tempList
	}

	d.Rows = 2
	d.Cols = len(d.Dates)
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
			tempRow[j] = strconv.FormatFloat(float64(value[j-1]), 'f', 3, 64)
		}
		data = append(data, tempRow)
	}

	file, err := os.Create(d.Target + "(" + d.Metric + ")" + ".csv")
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

type CompareDownloadData struct {
	Data1 []float32 //第一个仓库的数据
	Data2 []float32 //第二个仓库的数据
}

type CompareDownloadService struct {
	Source1 string
	Source2 string
	Target  string
	Title1  string
	Title2  string
	Dates   []string
	Data    map[string]CompareDownloadData
}

func (d *CompareDownloadService) SetData(source1_ RepoInfo, source2_ RepoInfo, target_ string) error {
	d.Target = target_
	d.Source1 = source1_.repoUrl
	d.Source2 = source2_.repoUrl
	d.Title1 = source1_.repoName
	d.Title2 = source2_.repoName
	d.Data = make(map[string]CompareDownloadData, 0)

	//dates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	//d.Dates = append(d.Dates, dates...)
	if len(source1_.dates) >= len(source2_.dates) {
		d.Dates = source1_.dates
	} else {
		d.Dates = source2_.dates
	}

	for k, v1 := range source1_.data {
		v2, ok := source2_.data[k]
		if ok {
			c := &CompareDownloadData{}
			data1 := make([]float32, 0)
			data2 := make([]float32, 0)

			for _, v3 := range d.Dates {
				temp1, ok1 := v1[v3]
				if ok1 {
					data1 = append(data1, parseFloatValue(temp1))
				} else {
					data1 = append(data1, 0)
				}
				temp2, ok2 := v2[v3]
				if ok2 {
					data2 = append(data2, parseFloatValue(temp2))
				} else {
					data2 = append(data2, 0)
				}
			}

			c.Data1 = data1
			c.Data2 = data2
			d.Data[k] = *c
		} else {
			continue
		}
	}
	return nil
}

func (d *CompareDownloadService) Download() error {

	tmpl, err := template.ParseFiles("../assets/template/template_compare.html")
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
