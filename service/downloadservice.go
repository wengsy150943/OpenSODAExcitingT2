package service

import (
	"encoding/csv"
	"errors"
	"exciting-opendigger/utils"
	"fmt"
	"go/types"
	"html/template"
	"log"
	"math"
	"os"
	"path"
	"sort"
	"strconv"
	"time"
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

var SpecialMetricForDownload = map[string]bool{
	"activity_details":                   true,
	"bus_factor_detail":                  true,
	"new_contributors_detail":            true,
	"active_dates_and_times":             true,
	"issue_response_time":                true,
	"issue_resolution_duration":          true,
	"change_request_response_time":       true,
	"change_request_resolution_duration": true,
	"change_request_age":                 true,
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

// 定义多折线图结构
type RaceLineData struct {
	RaceDates []string
	Avg       []float64
	Quantile0 []float64
	Quantile1 []float64
	Quantile2 []float64
	Quantile3 []float64
	Quantile4 []float64
}

func parseFloatValue(v interface{}) float32 {
	switch v.(type) {
	case float32:
		return v.(float32)
	case float64:
		return float32(v.(float64))
	case int:
		return float32(v.(int))
	case types.Nil:
		return 0.0
	}

	return 0.0
}

type UserDownloadService struct {
	User   string
	Source string
	Target string
	Dates  []string
	Data   map[string]([]float32)
}

func (d *UserDownloadService) SetData(target_ string, user_ string) error {
	userInfo := GetCertainUser(user_)
	data := utils.Usermetric{}
	data = utils.Parseuser(userInfo.Data, data)
	d.Data = make(map[string]([]float32))
	d.User = user_
	d.Source = "https://github.com/" + user_
	d.Target = target_
	d.Dates = userInfo.Dates

	//fmt.Println(data.Openrank)
	//fmt.Println(data.Activity)
	//fmt.Println(d.Dates)

	tempList1 := make([]float32, 0)
	for _, v2 := range d.Dates {
		temp, ok := data.Openrank[v2]
		if ok {
			tempList1 = append(tempList1, parseFloatValue(temp))
		} else {
			tempList1 = append(tempList1, 0)
		}
	}

	d.Data["openrank"] = tempList1

	tempList2 := make([]float32, 0)
	for _, v2 := range d.Dates {
		temp, ok := data.Activity[v2]
		if ok {
			tempList2 = append(tempList2, parseFloatValue(temp))
		} else {
			tempList2 = append(tempList2, 0)
		}
	}

	d.Data["activity"] = tempList2

	//fmt.Println(len(d.Data["activity"]))
	//fmt.Println(len(d.Data["openrank"]))
	//fmt.Println(len(d.Dates))

	return nil
}

func (d *UserDownloadService) Download() error {

	tmpl, err := template.ParseFiles("./assets/template/template_user.html")
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
	ActiveDatesAndTimesData   map[string]int
	QuantileStatsData         map[string]RaceLineData //负责五个分位数均值的metric

	//以下时单月查询使用
	MapDataOne map[string]float32
	Year       int
	Month      int
}

func (d *SingleDownloadService) SetData(source_ RepoInfo, target_ string) error {
	if source_.Dates[0] == "" || source_.Data == nil {
		return nil
	}

	d.Target = target_
	d.Source = source_.RepoUrl
	d.Title = source_.RepoName
	//fmt.Println(d.Title)
	d.Dates = source_.Dates
	d.Data = make(map[string]([]float32))
	d.QuantileStatsData = make(map[string]RaceLineData)

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

	for k, v := range source_.Data {
		if k == "active_dates_and_times" {
			newActiveDatesAndTimes := make(map[string]int)
			newActiveDatesAndTimes["test"] = 0
			d.ActiveDatesAndTimesData = newActiveDatesAndTimes
			//activeDatesAndTimes, years := getCalendarData(source_.SpecialData.ActiveDatesAndTimes)
			//d.ActiveDatesAndTimesData = activeDatesAndTimes
			//d.Years = getUnionOfTwoLists(d.Years, years)
			//fmt.Println("active")
			//fmt.Println(d.ActiveDatesAndTimes)
			//fmt.Println(d.Years)
		} else if k == "new_contributors_detail" {
			tempDetail := source_.SpecialData.NewContributorsDetail
			tempDetail2 := make(map[string]([][]string))
			for k1, v1 := range tempDetail {

				temp := make([][]string, 0)

				for _, v2 := range v1 {
					temp = append(temp, []string{v2, "1"})
				}

				tempDetail2[k1] = temp
			}
			w := getWordCloudData(tempDetail2)
			d.NewContributorsDetailData = w.YearData
			d.Years = getUnionOfTwoLists(d.Years, w.Years)
			//fmt.Println("new")
			//fmt.Println(d.NewContributorsDetailData)
			//fmt.Println(d.Years)
		} else if k == "bus_factor_detail" {
			w := getWordCloudData(source_.SpecialData.BusFactorDetail)
			d.BusFactorDetailData = w.YearData
			d.Years = getUnionOfTwoLists(d.Years, w.Years)
			//d.Years = w.Years
			//fmt.Println("bus")
			//fmt.Println(d.BusFactorDetailData)
			//fmt.Println(d.Years)
		} else if k == "activity_details" {
			w := getWordCloudData(source_.SpecialData.ActivityDetails)
			d.ActivityDetailsData = w.YearData
			d.Years = getUnionOfTwoLists(d.Years, w.Years)
			//d.Years = w.Years
			//fmt.Println("act")
			//fmt.Println(d.ActivityDetailsData)
			//fmt.Println(d.Years)
		} else if k == "issue_response_time" {
			d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.IssueResponseTime)
		} else if k == "issue_resolution_duration" {
			d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.IssueResolutionDuration)
		} else if k == "change_request_response_time" {
			d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.ChangeRequestResponseTime)
		} else if k == "change_request_resolution_duration" {
			d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.ChangeRequestResolutionDuration)
		} else if k == "change_request_age" {
			d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.ChangeRequestAge)
		} else if k == "issue_age" {
			d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.IssueAge)
		} else {
			tempList := make([]float32, 0)
			for _, v2 := range d.Dates {
				temp, ok := v[v2]
				if ok {
					tempList = append(tempList, parseFloatValue(temp))
				} else {
					tempList = append(tempList, 0)
				}
			}
			d.Data[k] = tempList
		}
	}

	return nil
}

func (d *SingleDownloadService) SetDataOneMetric(source_ RepoInfo, target_ string, k string) error {
	//保存在数据库中的空数据是"",从url获得的空数据是nil
	if source_.Dates[0] == "" || source_.Data == nil {
		return nil
	}
	d.Target = target_
	d.Source = source_.RepoUrl
	d.Title = source_.RepoName
	//fmt.Println(d.Title)
	d.Dates = source_.Dates
	d.Data = make(map[string]([]float32))
	d.QuantileStatsData = make(map[string]RaceLineData)
	initYear := 0
	initMonth := 0
	if d.Dates != nil {
		initYear, _ = strconv.Atoi(d.Dates[0][:4])
		initMonth, _ = strconv.Atoi(d.Dates[0][5:7])
	}

	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//
	//if err2 != nil {
	//	fmt.Println(err2)
	//}

	d.InitYear = initYear
	d.InitMonth = initMonth

	v1, err := source_.Data[k]

	if !err {
		return errors.New("invalid metric")
	}

	if k == "active_dates_and_times" {
		newActiveDatesAndTimes := make(map[string]int)
		newActiveDatesAndTimes["test"] = 0
		d.ActiveDatesAndTimesData = newActiveDatesAndTimes
		//activeDatesAndTimes, years := getCalendarData(source_.SpecialData.ActiveDatesAndTimes)
		//d.ActiveDatesAndTimesData = activeDatesAndTimes
		//d.Years = getUnionOfTwoLists(d.Years, years)
		//fmt.Println("active")
		//fmt.Println(d.ActiveDatesAndTimes)
		//fmt.Println(d.Years)
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
		d.Years = getUnionOfTwoLists(d.Years, w.Years)
		//fmt.Println("new")
		//fmt.Println(d.NewContributorsDetailData)
		//fmt.Println(d.Years)
	} else if k == "bus_factor_detail" {
		w := getWordCloudData(source_.SpecialData.BusFactorDetail)
		d.BusFactorDetailData = w.YearData
		d.Years = getUnionOfTwoLists(d.Years, w.Years)
		//d.Years = w.Years
		//fmt.Println("bus")
		//fmt.Println(d.BusFactorDetailData)
		//fmt.Println(d.Years)
	} else if k == "activity_details" {
		w := getWordCloudData(source_.SpecialData.ActivityDetails)
		d.ActivityDetailsData = w.YearData
		d.Years = getUnionOfTwoLists(d.Years, w.Years)
		//d.Years = w.Years
		//fmt.Println("act")
		//fmt.Println(d.ActivityDetailsData)
		//fmt.Println(d.Years)
	} else if k == "issue_response_time" {
		d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.IssueResponseTime)
	} else if k == "issue_resolution_duration" {
		d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.IssueResponseTime)
	} else if k == "change_request_response_time" {
		d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.IssueResponseTime)
	} else if k == "change_request_resolution_duration" {
		d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.IssueResponseTime)
	} else if k == "change_request_age" {
		d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.IssueResponseTime)
	} else if k == "issue_age" {
		d.QuantileStatsData[k] = getRaceLineData(source_.SpecialData.IssueAge)
	} else {
		tempList := make([]float32, 0)
		for _, v2 := range d.Dates {
			temp, ok := v1[v2]
			if ok {
				tempList = append(tempList, parseFloatValue(temp))
			} else {
				tempList = append(tempList, 0)
			}
		}
		d.Data[k] = tempList
	}

	return nil
}

func (d *SingleDownloadService) SetDataOneMonth(source_ RepoInfo, target_ string, year_ int, month_ int, metric_ string) error {
	//保存在数据库中的空数据是"",从url获得的空数据是nil
	if source_.Dates[0] == "" || source_.Data == nil {
		return nil
	}
	d.Target = target_
	d.Source = source_.RepoUrl
	d.Title = source_.RepoName
	//fmt.Println(d.Title)

	d.Dates = source_.Dates
	d.MapDataOne = make(map[string]float32)
	d.Years = []int{year_}
	d.InitYear = year_
	d.InitMonth = month_
	d.QuantileStatsData = make(map[string]RaceLineData)

	valid := false

	yearStr := strconv.Itoa(year_)
	monthStr := strconv.Itoa(month_)
	if len(monthStr) < 2 {
		monthStr = "0" + monthStr
	}

	date := yearStr + "-" + monthStr

	//fmt.Println(date)
	//
	//fmt.Println(d.Dates)

	for _, x := range d.Dates {
		if x == date {
			valid = true
		}
	}

	if valid == false {
		return errors.New("invalid date or no data ")
	}

	for k, v := range source_.Data {
		if k == "active_dates_and_times" {

			if len(metric_) != 0 && metric_ != k {
				continue
			}

			_, err := source_.SpecialData.ActiveDatesAndTimes[date]
			if !err {
				continue
			}

			newActiveDatesAndTimes := make(map[string]int)
			newActiveDatesAndTimes["test"] = 0
			d.ActiveDatesAndTimesData = newActiveDatesAndTimes

		} else if k == "new_contributors_detail" {
			if len(metric_) != 0 && metric_ != k {
				continue
			}
			tempDetail := source_.SpecialData.NewContributorsDetail
			tempDetail2 := make(map[string]([][]string))
			value, err := tempDetail[date]

			if !err {
				continue
			}
			temp := make([][]string, 0)

			for _, v1 := range value {
				temp = append(temp, []string{v1, "1"})
			}
			tempDetail2[date] = temp
			w := getWordCloudData(tempDetail2)
			d.NewContributorsDetailData = w.YearData

		} else if k == "bus_factor_detail" {
			if len(metric_) != 0 && metric_ != k {
				continue
			}

			tempDetail := source_.SpecialData.BusFactorDetail
			tempDetail2 := make(map[string]([][]string))
			tempDetail2[date] = tempDetail[date]
			w := getWordCloudData(tempDetail2)
			d.BusFactorDetailData = w.YearData

		} else if k == "activity_details" {
			if len(metric_) != 0 && metric_ != k {
				continue
			}
			tempDetail := source_.SpecialData.ActivityDetails
			tempDetail2 := make(map[string]([][]string))
			tempDetail2[date] = tempDetail[date]
			w := getWordCloudData(tempDetail2)
			d.ActivityDetailsData = w.YearData

		} else if k == "issue_response_time" {
			if len(metric_) != 0 && metric_ != k {
				continue
			}
			tempDetail := source_.SpecialData.IssueResponseTime
			tempDetail2 := make(map[string]utils.QuantileStats)
			tempDetail2[date] = tempDetail[date]
			d.QuantileStatsData[k] = getRaceLineData(tempDetail2)
		} else if k == "issue_resolution_duration" {
			if len(metric_) != 0 && metric_ != k {
				continue
			}
			tempDetail := source_.SpecialData.IssueResolutionDuration
			tempDetail2 := make(map[string]utils.QuantileStats)
			tempDetail2[date] = tempDetail[date]
			d.QuantileStatsData[k] = getRaceLineData(tempDetail2)
		} else if k == "change_request_response_time" {
			if len(metric_) != 0 && metric_ != k {
				continue
			}
			tempDetail := source_.SpecialData.ChangeRequestResponseTime
			tempDetail2 := make(map[string]utils.QuantileStats)
			tempDetail2[date] = tempDetail[date]
			d.QuantileStatsData[k] = getRaceLineData(tempDetail2)
		} else if k == "change_request_resolution_duration" {
			if len(metric_) != 0 && metric_ != k {
				continue
			}
			tempDetail := source_.SpecialData.ChangeRequestResolutionDuration
			tempDetail2 := make(map[string]utils.QuantileStats)
			tempDetail2[date] = tempDetail[date]
			d.QuantileStatsData[k] = getRaceLineData(tempDetail2)
		} else if k == "change_request_age" {
			if len(metric_) != 0 && metric_ != k {
				continue
			}
			tempDetail := source_.SpecialData.ChangeRequestAge
			tempDetail2 := make(map[string]utils.QuantileStats)
			tempDetail2[date] = tempDetail[date]
			d.QuantileStatsData[k] = getRaceLineData(tempDetail2)
		} else if k == "issue_age" {
			if len(metric_) != 0 && metric_ != k {
				continue
			}
			tempDetail := source_.SpecialData.IssueAge
			tempDetail2 := make(map[string]utils.QuantileStats)
			tempDetail2[date] = tempDetail[date]
			d.QuantileStatsData[k] = getRaceLineData(tempDetail2)
		} else {
			if len(metric_) != 0 && metric_ != k {
				continue
			}
			if v[date] == nil {
				println(v[date])
			}
			d.MapDataOne[k] = parseFloatValue(v[date])
		}
	}

	//fmt.Println(d.MapDataOne)

	return nil
}

func (d *SingleDownloadService) Download() error {

	tmpl, err := template.ParseFiles("./assets/template/template.html")
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

func (d *SingleDownloadService) DownloadMonth() error {

	tmpl, err := template.ParseFiles("./assets/template/template_month.html")
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
	Rows   int
	Cols   int
	Data   map[string]([]interface{}) //这里的key为仓库名
	Dates  []string
}

func pathExists(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		if e := os.Mkdir(path, os.ModePerm); e != nil {
			return e
		}
		return nil
	}
	if !s.IsDir() {
		return errors.New("the path point to a file, not a dir")
	}
	return nil
}

func (d *BatchDownloadService) SetData(sources_ []RepoInfo, metric_ string) error {
	d.Metric = metric_
	maxLength := 0
	for _, repo := range sources_ {
		if len(repo.Dates) > maxLength {
			d.Dates = repo.Dates
			maxLength = len(d.Dates)
		}
	}

	if SpecialMetricForDownload[metric_] == true {
		return errors.New("unsupported metric")
	}

	d.Data = make(map[string]([]interface{}))

	for _, repo := range sources_ {
		name := repo.RepoName
		tempList := make([]interface{}, 0)
		for _, v := range d.Dates {
			temp, ok := repo.Data[metric_][v]
			if ok {
				tempList = append(tempList, temp)
			} else {
				tempList = append(tempList, 0)
			}
		}
		d.Data[name] = tempList
	}

	d.Rows = len(sources_)
	d.Cols = len(d.Dates)
	return nil
}

func (d *BatchDownloadService) Download(filepath string) error {
	if err := pathExists(filepath); err != nil {
		log.Fatal(err)
	}
	data := make([][]string, 0)

	firstRow := make([]string, 0)
	firstRow = append(firstRow, "RepoName")
	firstRow = append(firstRow, d.Dates...)

	data = append(data, firstRow)

	for key, value := range d.Data {
		tempRow := make([]string, d.Cols+1)
		tempRow[0] = key
		for j := 1; j <= d.Cols; j++ {
			tempRow[j] = fmt.Sprintf("%v", value[j-1])
		}
		data = append(data, tempRow)
	}

	file, err := os.Create(path.Join(filepath, d.Metric+".csv"))
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

	if len(source1_.Dates) >= len(source2_.Dates) {
		d.Dates = source1_.Dates
	} else {
		d.Dates = source2_.Dates
	}

	for k, v1 := range source1_.Data {

		if SpecialMetricForDownload[k] == true {
			continue
		}

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

	tmpl, err := template.ParseFiles("./assets/template/template_compare.html")
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

func getUnionOfTwoLists(listA_ []int, listB_ []int) []int {
	listA := listA_
	listB := listB_

	union := make(map[int]bool)

	// 将列表 A 中的元素添加到并集
	for _, num := range listA {
		union[num] = true
	}

	// 将列表 B 中的元素添加到并集
	for _, num := range listB {
		union[num] = true
	}

	// 将并集转换回切片
	result := make([]int, 0, len(union))
	for num := range union {
		result = append(result, num)
	}

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	return result
}

func getCalendarData(data_ map[string]([]int)) (map[string]int, []int) {

	union := make(map[int]bool)
	res := make(map[string]int)
	for k, v := range data_ {

		//fmt.Println(k)
		//fmt.Println(v)
		//fmt.Println(len(v))

		yearMonth, err := time.Parse("2006-01", k)
		if err != nil {
			fmt.Println("Invalid date format:", k)
			continue
		}

		// 获取该月的天数
		year, month, _ := yearMonth.Date()
		union[year] = true
		_, _, lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Date()

		if lastDay != len(v) {
			//fmt.Println("Invalid month data:", k, " ", lastDay, " ", len(v))
			continue
		}

		// 生成年份-月份-日列表
		for day := 1; day <= lastDay; day++ {
			date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			//fmt.Println(date.Format("2006-01-02"))
			res[date.Format("2006-01-02")] = v[day-1]
		}
	}

	res2 := make([]int, 0, len(union))
	for num := range union {
		res2 = append(res2, num)
	}
	return res, res2
}

func getRaceLineData(data_ map[string]utils.QuantileStats) RaceLineData {
	var raceDates []string
	var avg []float64
	var quantile0 []float64
	var quantile1 []float64
	var quantile2 []float64
	var quantile3 []float64
	var quantile4 []float64

	dates := make([]string, 0)
	for k, _ := range data_ {
		dates = append(dates, k)
	}
	sort.Slice(dates, func(i, j int) bool { return dates[i] < dates[j] })

	for _, k := range dates {
		v := data_[k]
		raceDates = append(raceDates, k)
		avg = append(avg, v.Avg)
		quantile0 = append(quantile0, v.Quantile[0])
		quantile1 = append(quantile1, v.Quantile[1])
		quantile2 = append(quantile2, v.Quantile[2])
		quantile3 = append(quantile3, v.Quantile[3])
		quantile4 = append(quantile4, v.Quantile[4])
	}

	return RaceLineData{Avg: avg, RaceDates: raceDates, Quantile0: quantile0, Quantile1: quantile1, Quantile2: quantile2, Quantile3: quantile3, Quantile4: quantile4}
}
