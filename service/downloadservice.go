package service

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
	"strconv"
)

var MonthMap = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

type WordCloudData struct {
	YearData []YearMonthData
	Years    []int
}

// 定义云图结构
type WordCloudDetailData struct {
	Name  string
	Value float32
}

// 定义年份和月份的数据结构
type YearMonthData struct {
	Year  int
	Month int
	Data  []WordCloudDetailData
}

type SingleDownloadService struct {
	Source                    string
	Target                    string
	Title                     string
	Dates                     []string
	Data                      map[string]([]float32)
	Years                     []int
	InitYear                  int //前端按钮默认显示
	InitMonth                 int //前端按钮默认显示
	ActivityDetailsData       []YearMonthData
	BusFactorDetailData       []YearMonthData
	NewContributorsDetailData []YearMonthData
}

func parseFloatValue(v interface{}) float32 {
	switch v.(type) {
	case float32:
		return v.(float32)
	case float64:
		return float32(v.(float64))
	case int:
		return float32(v.(int))
	}
	return v.(float32)
}

func (d *SingleDownloadService) SetData(source_ RepoInfo, target_ string) error {
	d.Target = target_
	d.Source = source_.RepoUrl
	d.Title = source_.RepoName

	d.Dates = source_.Dates
	d.Data = make(map[string]([]float32))

	initYear, err1 := strconv.Atoi(d.Dates[0][:4])
	initMonth, err2 := strconv.Atoi(d.Dates[0][5:7])

	if err1 != nil {
		fmt.Println(err1)
	}

	if err2 != nil {
		fmt.Println(err2)
	}

	d.InitYear = initYear
	d.InitMonth = initMonth

	for k, v1 := range source_.Data {
		if k == "active_dates_and_times" {
			continue
		} else if k == "new_contributors_detail" {
			tempDetail := source_.SpecialData.NewContributorsDetail
			tempDetail2 := make(map[string]([][]string))
			for k, v := range tempDetail {

				temp := make([][]string, 0)

				for _, v1 := range v {
					temp = append(temp, []string{v1, "1"})
				}

				tempDetail2[k] = temp
			}
			w := getWordCloudData(tempDetail2)
			d.NewContributorsDetailData = w.YearData
			d.Years = w.Years
			fmt.Println("new")
			fmt.Println(d.NewContributorsDetailData)
			fmt.Println(d.Years)
		} else if k == "bus_factor_detail" {
			w := getWordCloudData(source_.SpecialData.BusFactorDetail)
			d.BusFactorDetailData = w.YearData
			d.Years = w.Years
			fmt.Println("bus")
			fmt.Println(d.BusFactorDetailData)
			fmt.Println(d.Years)
		} else if k == "activity_details" {
			w := getWordCloudData(source_.SpecialData.ActivityDetails)
			d.ActivityDetailsData = w.YearData
			d.Years = w.Years
			fmt.Println("act")
			fmt.Println(d.ActivityDetailsData)
			fmt.Println(d.Years)
		} else if k == "issue_response_time" {
			continue
		} else if k == "issue_resolution_duration" {
			continue
		} else if k == "change_request_response_time" {
			continue
		} else if k == "change_request_resolution_duration" {
			continue
		} else if k == "change_request_age" {
			continue
		} else {
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

	Dates []string
}

func (d *BatchDownloadService) SetData(sources_ []RepoInfo, metric_ string, target_ string) error {
	d.Target = target_
	d.Metric = metric_
	maxLength := 0
	for _, repo := range sources_ {
		if len(repo.Dates) > maxLength {
			d.Dates = repo.Dates
			maxLength = len(d.Dates)
		}
	}

	d.Data = make(map[string]([]float32))

	for _, repo := range sources_ {
		name := repo.RepoName
		tempList := make([]float32, 0)
		for _, v := range d.Dates {
			temp, ok := repo.Data[metric_][v]
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
	d.Source1 = source1_.RepoUrl
	d.Source2 = source2_.RepoUrl
	d.Title1 = source1_.RepoName
	d.Title2 = source2_.RepoName
	d.Data = make(map[string]CompareDownloadData, 0)

	//dates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	//d.Dates = append(d.Dates, dates...)
	if len(source1_.Dates) >= len(source2_.Dates) {
		d.Dates = source1_.Dates
	} else {
		d.Dates = source2_.Dates
	}

	for k, v1 := range source1_.Data {
		v2, ok := source2_.Data[k]
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

func getWordCloudData(data_ map[string]([][]string)) WordCloudData {
	yearMin := math.MaxInt
	yearMax := math.MinInt
	wordCloudData := make(map[int]map[int][]WordCloudDetailData)

	for key, item := range data_ {

		year, err := strconv.Atoi(key[:4])
		if year < yearMin {
			yearMin = year
		}

		if year > yearMax {
			yearMax = year
		}

		if wordCloudData[year] == nil {
			wordCloudData[year] = make(map[int][]WordCloudDetailData)
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		month, err1 := strconv.Atoi(key[5:7])
		if err != nil {
			fmt.Println(err1)
			continue
		}

		for _, value := range item {
			score, err2 := strconv.ParseFloat(value[1], 32)
			if err2 != nil {
				fmt.Println(err2)
				continue
			}
			temp := &WordCloudDetailData{value[0], float32(score)}

			wordCloudData[year][month] = append(wordCloudData[year][month], *temp)
		}

	}

	//fmt.Println(data_)

	// 创建一个用于存储年份和月份数据的切片
	var yearMonthData []YearMonthData

	// 转换为 YearMonthData 结构，并添加到 yearMonthData 切片中
	for year, monthData := range wordCloudData {
		for month, data := range monthData {
			yearMonthData = append(yearMonthData, YearMonthData{
				Year:  year,
				Month: month,
				Data:  data,
			})
		}
	}

	// 计算数组的长度
	length := yearMax - yearMin + 1

	// 创建并初始化连续数组
	years := make([]int, length)
	for i := 0; i < length; i++ {
		years[i] = yearMin + i
	}

	res := &WordCloudData{yearMonthData, years}

	return *res

}
