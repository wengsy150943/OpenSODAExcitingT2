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
	Active_dates_and_times  map[string]([]int)      //key为日期
	New_contributors_detail map[string]([]string)   //key为日期
	Bus_factor_detail       map[string]([][]string) //key为日期，value这边的float可以由string转
	Activity_details        map[string]([][]string) //key为日期，value这边的float可以由string转
	// 以下部分都是转为QuantileStats格式的数据
	Issue_response_time                map[string](QuantileStats) //key为日期
	Issue_resolution_duration          map[string](QuantileStats)
	Change_request_response_time       map[string](QuantileStats)
	Change_request_resolution_duration map[string](QuantileStats)
	Change_request_age                 map[string](QuantileStats)
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

// 用反射遍历所有变量，逐个暴力枚举把月份清理掉
func (r *SpecialDataStructure) SelectMonth(month string) {

	value := reflect.ValueOf(r).Elem()
	dataType := reflect.TypeOf(r).Elem()
	for i := 0; i < value.NumField(); i++ {
		if value.FieldByName(dataType.Name()).IsValid() {
			val := value.FieldByName(dataType.Name())
			dataMap := reflect.MakeMap(val.Type())

			iter := val.MapRange()
			for iter.Next() {
				keyValue := iter.Key()
				value := iter.Value()
				// 只设置对的月份
				if keyValue.String() == month {
					dataMap.SetMapIndex(keyValue, value)
				}
			}
			// 更新值
			val.Set(dataMap)
		}
	}
}

func parseActiveDatesAndTimes(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.Active_dates_and_times = make(map[string][]int)
	for month, v := range data {
		valueList := make([]int, len(v.([]interface{})))
		for key, value := range v.([]interface{}) {
			valueList[key] = int(value.(float64))
		}
		r.Active_dates_and_times[month] = valueList
	}
	return r
}

func parseNewContributorsDetail(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.New_contributors_detail = make(map[string][]string)
	for month, v := range data {
		valueList := make([]string, len(v.([]interface{})))
		for key, value := range v.([]interface{}) {
			valueList[key] = value.(string)
		}
		r.New_contributors_detail[month] = valueList
	}
	return r
}

// 这两个的内层都是[string, float],为了组织方便强转为[]string
func parseBusFactorDetail(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.Bus_factor_detail = make(map[string][][]string)
	for month, v := range data {

		val := v.([]interface{})
		valueList := make([][]string, len(val))
		for key, value := range val {
			temp := make([]string, 2)
			temp[0] = value.([]interface{})[0].(string)
			temp[1] = strconv.FormatFloat(value.([]interface{})[1].(float64), 'f', -1, 64)
			valueList[key] = temp
		}
		r.Bus_factor_detail[month] = valueList
	}
	return r
}

func parseActivityDetails(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.Activity_details = make(map[string][][]string)
	for month, v := range data {
		val := v.([]interface{})
		valueList := make([][]string, len(val))
		for key, value := range val {
			temp := make([]string, 2)
			temp[0] = value.([]interface{})[0].(string)
			temp[1] = strconv.FormatFloat(value.([]interface{})[1].(float64), 'f', -1, 64)
			valueList[key] = temp
		}
		r.Activity_details[month] = valueList
	}
	return r
}

//以下函数均是将对应数据类型转换为QuantileStats的map
// 参数：
//   data: 对应metric的数据项

func parseIssueResponseTime(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.Issue_response_time = make(map[string]QuantileStats)

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
		r.Issue_response_time[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}

func parseIssueResolutionDuration(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.Issue_resolution_duration = make(map[string]QuantileStats)

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
		r.Issue_resolution_duration[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}

func parseChangeRequestResponseTime(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {

	r.Change_request_response_time = make(map[string]QuantileStats)

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
		r.Change_request_response_time[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}

func parseChangeRequestResolutionDuration(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.Change_request_resolution_duration = make(map[string]QuantileStats)

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
		r.Change_request_resolution_duration[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}

func parseChangeRequestAge(data map[string]interface{}, r SpecialDataStructure) SpecialDataStructure {
	r.Change_request_age = make(map[string]QuantileStats)

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
		r.Change_request_age[month] = QuantileStats{
			Levels:   levels,
			Quantile: quantile,
			Avg:      data["avg"].(map[string]interface{})[month].(float64),
		}
	}
	return r
}
