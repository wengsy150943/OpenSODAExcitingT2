package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// 这类数据包括三部分，levels int列表，quantile 四分位数 这里统一用一个float列表存， avg 平均值 float
type QuantileStats struct {
	Levels   []int
	Quantile []float64 // quantile 0,1,2,3,4
	Avg      float64
}

type SpecialDataStructure struct {
	ActiveDatesAndTimes   map[string]([]int)      //key为日期
	NewContributorsDetail map[string]([]string)   //key为日期
	BusFactorDetail       map[string]([][]string) //key为日期，value这边的float可以由string转
	ActivityDetails       map[string]([][]string) //key为日期，value这边的float可以由string转
	// 以下部分都是转为QuantileStats格式的数据
	IssueResponseTime               map[string](QuantileStats) //key为日期
	IssueResolutionDuration         map[string](QuantileStats)
	ChangeRequestResponseTime       map[string](QuantileStats)
	ChangeRequestResolutionDuration map[string](QuantileStats)
	ChangeRequestAge                map[string](QuantileStats)
}

type parseFunc func(map[string]interface{}, SpecialDataStructure) SpecialDataStructure

// 用函数map存所有的函数，后续使用起来可以用tag直接调
var Parse = map[string]parseFunc{
	"active_dates_and_times":             parseActiveDatesAndTimes,
	"new_contributors_detail":            parseNewContributorsDetail,
	"bus_factor_detail":                  parseBusFactorDetail,
	"activity_details":                   parseActivityDetails,
	"issue_response_time":                parseIssueResponseTime,
	"issue_resolution_duration":          parseIssueResolutionDuration,
	"change_request_response_time":       parseChangeRequestResponseTime,
	"change_request_resolution_duration": parseChangeRequestResolutionDuration,
	"change_request_age":                 parseChangeRequestAge,
}

func (r *SpecialDataStructure) MergeSpecialData(others SpecialDataStructure) {
	value := reflect.ValueOf(r).Elem()
	dataType := reflect.TypeOf(r).Elem()

	otherValue := reflect.ValueOf(others)
	for i := 0; i < value.NumField(); i++ {
		if otherValue.FieldByName(dataType.Field(i).Name).IsValid() {
			val := reflect.New(dataType.Field(i).Type)
			val.Elem().Set(otherValue.FieldByName(dataType.Field(i).Name))

			if val.Elem().Len() > 0 {
				value.Field(i).Set(val.Elem())
			}

		}
	}
}

// 用反射遍历所有变量，逐个暴力枚举把月份清理掉
func (r *SpecialDataStructure) SelectMonth(month string) {

	value := reflect.ValueOf(r).Elem()
	dataType := reflect.TypeOf(r).Elem()
	for i := 0; i < dataType.NumField(); i++ {

		val := value.FieldByName(dataType.Field(i).Name)
		if val.IsValid() && len(val.MapKeys()) > 0 {

			iter := val.MapRange()
			for iter.Next() {
				keyValue := iter.Key()
				// 只设置对的月份
				if keyValue.String() != month {
					val.SetMapIndex(keyValue, reflect.Value{})
				}
			}
		}
	}
}

func parseActiveDatesAndTimes(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.ActiveDatesAndTimes = make(map[string][]int)
	for month, v := range data {
		valueList := make([]int, len(v.([]interface{})))
		for key, value := range v.([]interface{}) {
			valueList[key] = int(value.(float64))
		}
		r.ActiveDatesAndTimes[month] = valueList
	}
	return r
}

func parseNewContributorsDetail(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.NewContributorsDetail = make(map[string][]string)
	for month, v := range data {
		valueList := make([]string, len(v.([]interface{})))
		for key, value := range v.([]interface{}) {
			valueList[key] = value.(string)
		}
		r.NewContributorsDetail[month] = valueList
	}
	return r
}

// 这两个的内层都是[string, float],为了组织方便强转为[]string
func parseBusFactorDetail(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.BusFactorDetail = make(map[string][][]string)
	for month, v := range data {

		val := v.([]interface{})
		valueList := make([][]string, len(val))
		for key, value := range val {
			temp := make([]string, 2)
			temp[0] = value.([]interface{})[0].(string)
			temp[1] = strconv.FormatFloat(value.([]interface{})[1].(float64), 'f', -1, 64)
			valueList[key] = temp
		}
		r.BusFactorDetail[month] = valueList
	}
	return r
}

func parseActivityDetails(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.ActivityDetails = make(map[string][][]string)
	for month, v := range data {
		val := v.([]interface{})
		valueList := make([][]string, len(val))
		for key, value := range val {
			temp := make([]string, 2)
			temp[0] = value.([]interface{})[0].(string)
			temp[1] = strconv.FormatFloat(value.([]interface{})[1].(float64), 'f', -1, 64)
			valueList[key] = temp
		}
		r.ActivityDetails[month] = valueList
	}
	return r
}

//以下函数均是将对应数据类型转换为QuantileStats的map
// 参数：
//   data: 对应metric的数据项

func parseIssueResponseTime(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.IssueResponseTime = make(map[string]QuantileStats)

	for month, _ := range data["avg"].(map[string]interface{}) {

		quantile := make([]float64, 5)
		for i := 0; i < 5; i++ {
			tag := "quantile_" + fmt.Sprint(i)
			quantile[i] = data[tag].(map[string]interface{})[month].(float64)
		}
		dataMap := data["levels"].(map[string](interface{}))
		levels := make([]int, len(dataMap[month].([]interface{})))
		for k, v := range dataMap[month].([]interface{}) {
			levels[k] = int(v.(float64))
		}
		r.IssueResponseTime[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}

func parseIssueResolutionDuration(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.IssueResolutionDuration = make(map[string]QuantileStats)

	for month, _ := range data["avg"].(map[string]interface{}) {

		quantile := make([]float64, 5)
		for i := 0; i < 5; i++ {
			tag := "quantile_" + fmt.Sprint(i)
			quantile[i] = data[tag].(map[string]interface{})[month].(float64)
		}
		dataMap := data["levels"].(map[string](interface{}))
		levels := make([]int, len(dataMap[month].([]interface{})))
		for k, v := range dataMap[month].([]interface{}) {
			levels[k] = int(v.(float64))
		}
		r.IssueResolutionDuration[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}

func parseChangeRequestResponseTime(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {

	r.ChangeRequestResponseTime = make(map[string]QuantileStats)

	for month, _ := range data["avg"].(map[string]interface{}) {

		quantile := make([]float64, 5)
		for i := 0; i < 5; i++ {
			tag := "quantile_" + fmt.Sprint(i)
			quantile[i] = data[tag].(map[string]interface{})[month].(float64)
		}
		dataMap := data["levels"].(map[string](interface{}))
		levels := make([]int, len(dataMap[month].([]interface{})))
		for k, v := range dataMap[month].([]interface{}) {
			levels[k] = int(v.(float64))
		}
		r.ChangeRequestResponseTime[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}

func parseChangeRequestResolutionDuration(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.ChangeRequestResolutionDuration = make(map[string]QuantileStats)

	for month, _ := range data["avg"].(map[string]interface{}) {

		quantile := make([]float64, 5)
		for i := 0; i < 5; i++ {
			tag := "quantile_" + fmt.Sprint(i)
			quantile[i] = data[tag].(map[string]interface{})[month].(float64)
		}
		dataMap := data["levels"].(map[string](interface{}))
		levels := make([]int, len(dataMap[month].([]interface{})))
		for k, v := range dataMap[month].([]interface{}) {
			levels[k] = int(v.(float64))
		}
		r.ChangeRequestResolutionDuration[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}

func parseChangeRequestAge(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.ChangeRequestAge = make(map[string]QuantileStats)

	for month, _ := range data["avg"].(map[string]interface{}) {

		quantile := make([]float64, 5)
		for i := 0; i < 5; i++ {
			tag := "quantile_" + fmt.Sprint(i)
			quantile[i] = data[tag].(map[string]interface{})[month].(float64)
		}
		dataMap := data["levels"].(map[string](interface{}))
		levels := make([]int, len(dataMap[month].([]interface{})))
		for k, v := range dataMap[month].([]interface{}) {
			levels[k] = int(v.(float64))
		}
		r.ChangeRequestAge[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}
