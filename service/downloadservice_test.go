package service

import (
	"exciting-opendigger/utils"
	"fmt"
	"testing"
)

func TestUserDownloadService(t *testing.T) {
	fmt.Println("TestUserDownloadService：")

	downloadService := &UserDownloadService{}
	err := downloadService.SetData("testUser", "will-ww")
	if err != nil {
		panic("fail to SetData")
	}
	err2 := downloadService.Download()
	if err2 != nil {
		panic("fail to Download")
	}
}

func TestSingleDownloadService(t *testing.T) {
	fmt.Println("TestSingleDownloadService：")
	downloadService := &SingleDownloadService{}
	testDates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testData2 := map[string]interface{}{"2020-08": 10, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testDataMap := make(map[string](map[string]interface{}))
	testDataMap["metricOne"] = testData1
	testDataMap["metricTwo"] = testData2

	testDataMap["bus_factor_detail"] = nil
	testDataMap["new_contributors_detail"] = nil
	testDataMap["activity_details"] = nil
	testDataMap["active_dates_and_times"] = nil

	testDataMap["issue_response_time"] = nil
	testDataMap["issue_resolution_duration"] = nil
	testDataMap["change_request_response_time"] = nil
	testDataMap["change_request_resolution_duration"] = nil
	testDataMap["change_request_age"] = nil

	myMap := make(map[string]([][]string))

	// 添加值到 map 中
	myMap["2020-08"] = [][]string{{"sunshinemingo1", "5"}, {"frank-zsy1", "12"}}
	myMap["2021-09"] = [][]string{{"sunshinemingo2", "50"}, {"frank-zsy2", "6"}}
	myMap["2022-10"] = [][]string{{"sunshinemingo3", "500"}, {"frank-zsy3", "120"}}

	myMap2 := make(map[string]([]string))

	// 添加值到 map 中
	myMap2["2020-08"] = []string{"sunshinemingo1new", "frank-zsy1new"}
	myMap2["2020-09"] = []string{"sunshinemingo2new", "frank-zsy2new"}
	myMap2["2020-10"] = []string{"sunshinemingo3new", "frank-zsy3new"}

	myMap3 := make(map[string]([]int))
	myMap3["2020-02"] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100}
	myMap3["2020-09"] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 50}
	myMap3["2021-02"] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	myMap3["2022-01"] = []int{1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 111, 11, 0}

	myMap4 := make(map[string]utils.QuantileStats)
	temp := utils.QuantileStats{Avg: 2, Quantile: []float64{1, 2, 3, 4, 5}, Levels: nil}

	myMap4["2020-02"] = temp
	myMap4["2020-09"] = temp
	myMap4["2021-02"] = temp
	myMap4["2022-01"] = temp

	testSpecial := &utils.SpecialDataStructure{BusFactorDetail: myMap, ActivityDetails: myMap, NewContributorsDetail: myMap2, ActiveDatesAndTimes: myMap3, IssueResponseTime: myMap4, IssueResolutionDuration: myMap4, ChangeRequestResolutionDuration: myMap4, ChangeRequestAge: myMap4, ChangeRequestResponseTime: myMap4}

	ret := RepoInfo{
		RepoName:    "opendigger",
		RepoUrl:     "www.github.com/X-lab2017/open-digger",
		Month:       "",
		Dates:       testDates,
		Data:        testDataMap,
		SpecialData: *testSpecial,
	}

	err := downloadService.SetData(ret, "html_output")
	if err != nil {
		return
	}
	err2 := downloadService.Download()
	if err2 != nil {
		return
	}
}

func TestSingleDownloadServiceOneMetric(t *testing.T) {
	fmt.Println("TestSingleDownloadServiceOneMetric：")

	testDates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testData2 := map[string]interface{}{"2020-08": 10, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testDataMap := make(map[string](map[string]interface{}))
	testDataMap["openrank"] = testData1
	testDataMap["openrank2"] = testData2

	testDataMap["bus_factor_detail"] = nil
	testDataMap["new_contributors_detail"] = nil
	testDataMap["activity_details"] = nil
	testDataMap["active_dates_and_times"] = nil

	testDataMap["issue_response_time"] = nil
	testDataMap["issue_resolution_duration"] = nil
	testDataMap["change_request_response_time"] = nil
	testDataMap["change_request_resolution_duration"] = nil
	testDataMap["change_request_age"] = nil

	myMap := make(map[string]([][]string))

	// 添加值到 map 中
	myMap["2020-08"] = [][]string{{"sunshinemingo1", "5"}, {"frank-zsy1", "12"}}
	myMap["2021-09"] = [][]string{{"sunshinemingo2", "50"}, {"frank-zsy2", "6"}}
	myMap["2022-10"] = [][]string{{"sunshinemingo3", "500"}, {"frank-zsy3", "120"}}

	myMap2 := make(map[string]([]string))

	// 添加值到 map 中
	myMap2["2020-08"] = []string{"sunshinemingo1new", "frank-zsy1new"}
	myMap2["2020-09"] = []string{"sunshinemingo2new", "frank-zsy2new"}
	myMap2["2020-10"] = []string{"sunshinemingo3new", "frank-zsy3new"}

	myMap3 := make(map[string]([]int))
	myMap3["2020-02"] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100}
	myMap3["2020-09"] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 50}
	myMap3["2021-02"] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	myMap3["2022-01"] = []int{1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 111, 11, 0}

	myMap4 := make(map[string]utils.QuantileStats)
	temp := utils.QuantileStats{Avg: 2, Quantile: []float64{1, 2, 3, 4, 5}, Levels: nil}

	myMap4["2020-02"] = temp
	myMap4["2020-09"] = temp
	myMap4["2021-02"] = temp
	myMap4["2022-01"] = temp

	testSpecial := &utils.SpecialDataStructure{BusFactorDetail: myMap, ActivityDetails: myMap, NewContributorsDetail: myMap2, ActiveDatesAndTimes: myMap3, IssueResponseTime: myMap4, IssueResolutionDuration: myMap4, ChangeRequestResolutionDuration: myMap4, ChangeRequestAge: myMap4, ChangeRequestResponseTime: myMap4}

	ret := RepoInfo{
		RepoName:    "opendigger",
		RepoUrl:     "www.github.com/X-lab2017/open-digger",
		Month:       "",
		Dates:       testDates,
		Data:        testDataMap,
		SpecialData: *testSpecial,
	}
	downloadService := &SingleDownloadService{}
	err := downloadService.SetDataOneMetric(ret, "html_output1", "openrank")
	if err != nil {
		fmt.Println(err)
	}

	err2 := downloadService.Download()
	if err2 != nil {
		fmt.Println(err2)
	}

	downloadService = &SingleDownloadService{}
	err3 := downloadService.SetDataOneMetric(ret, "html_output2", "issue_response_time")
	if err3 != nil {
		fmt.Println(err3)
	}
	err4 := downloadService.Download()
	if err4 != nil {
		fmt.Println(err4)
	}

	downloadService = &SingleDownloadService{}
	err5 := downloadService.SetDataOneMetric(ret, "html_output3", "active_dates_and_times")
	if err5 != nil {
		fmt.Println(err5)
	}
	err6 := downloadService.Download()
	if err6 != nil {
		fmt.Println(err6)
	}

	downloadService = &SingleDownloadService{}
	err7 := downloadService.SetDataOneMetric(ret, "html_output4", "bus_factor_detail")
	if err7 != nil {
		fmt.Println(err5)
	}
	err8 := downloadService.Download()
	if err8 != nil {
		fmt.Println(err8)
	}

}

func TestSingleDownloadServiceOneMonth(t *testing.T) {
	fmt.Println("TestSingleDownloadService：")

	testDates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testData2 := map[string]interface{}{"2020-08": 10, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testDataMap := make(map[string](map[string]interface{}))
	testDataMap["openrank"] = testData1
	testDataMap["metricTwo"] = testData2

	testDataMap["bus_factor_detail"] = nil
	testDataMap["new_contributors_detail"] = nil
	testDataMap["activity_details"] = nil
	testDataMap["active_dates_and_times"] = nil

	testDataMap["issue_response_time"] = nil
	testDataMap["issue_resolution_duration"] = nil
	testDataMap["change_request_response_time"] = nil
	testDataMap["change_request_resolution_duration"] = nil
	testDataMap["change_request_age"] = nil

	myMap := make(map[string]([][]string))

	// 添加值到 map 中
	myMap["2020-08"] = [][]string{{"sunshinemingo1", "5"}, {"frank-zsy1", "12"}}
	myMap["2021-09"] = [][]string{{"sunshinemingo2", "50"}, {"frank-zsy2", "6"}}
	myMap["2022-10"] = [][]string{{"sunshinemingo3", "500"}, {"frank-zsy3", "120"}}

	myMap2 := make(map[string]([]string))

	// 添加值到 map 中
	myMap2["2020-08"] = []string{"sunshinemingo1new", "frank-zsy1new"}
	myMap2["2020-09"] = []string{"sunshinemingo2new", "frank-zsy2new"}
	myMap2["2020-10"] = []string{"sunshinemingo3new", "frank-zsy3new"}

	myMap3 := make(map[string]([]int))
	myMap3["2020-08"] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 1}
	myMap3["2020-09"] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 50}
	myMap3["2021-02"] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	myMap3["2022-01"] = []int{1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 111, 11, 0}

	myMap4 := make(map[string]utils.QuantileStats)
	temp := utils.QuantileStats{Avg: 2, Quantile: []float64{1, 2, 3, 4, 5}, Levels: nil}

	myMap4["2020-08"] = temp
	myMap4["2020-09"] = temp
	myMap4["2021-02"] = temp
	myMap4["2022-01"] = temp

	testSpecial := &utils.SpecialDataStructure{BusFactorDetail: myMap, ActivityDetails: myMap, NewContributorsDetail: myMap2, ActiveDatesAndTimes: myMap3, IssueResponseTime: myMap4, IssueResolutionDuration: myMap4, ChangeRequestResolutionDuration: myMap4, ChangeRequestAge: myMap4, ChangeRequestResponseTime: myMap4}

	ret := RepoInfo{
		RepoName:    "opendigger",
		RepoUrl:     "www.github.com/X-lab2017/open-digger",
		Month:       "",
		Dates:       testDates,
		Data:        testDataMap,
		SpecialData: *testSpecial,
	}

	downloadService := &SingleDownloadService{}
	err := downloadService.SetDataOneMonth(ret, "202008", 2020, 8, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	err2 := downloadService.DownloadMonth()
	if err2 != nil {
		return
	}

	downloadService = &SingleDownloadService{}
	err = downloadService.SetDataOneMonth(ret, "202008_issue_openrank", 2020, 8, "openrank")
	if err != nil {
		return
	}
	err2 = downloadService.DownloadMonth()
	if err2 != nil {
		return
	}

	downloadService = &SingleDownloadService{}
	err = downloadService.SetDataOneMonth(ret, "202008_issue_response_time", 2020, 8, "issue_response_time")
	if err != nil {
		return
	}
	err2 = downloadService.DownloadMonth()
	if err2 != nil {
		return
	}

	downloadService = &SingleDownloadService{}
	err = downloadService.SetDataOneMonth(ret, "202008_issue_active_dates_and_times", 2020, 8, "active_dates_and_times")
	if err != nil {
		return
	}
	err2 = downloadService.DownloadMonth()
	if err2 != nil {
		return
	}

	downloadService = &SingleDownloadService{}
	err = downloadService.SetDataOneMonth(ret, "202008_issue_bus_factor_detail", 2020, 8, "bus_factor_detail")
	if err != nil {
		return
	}
	err2 = downloadService.DownloadMonth()
	if err2 != nil {
		return
	}

}

func TestBatchDownloadService(t *testing.T) {
	fmt.Println("TestBatchDownloadService：")
	downloadService := &BatchDownloadService{}

	var rets []RepoInfo

	testDates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testData2 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testDataMap := make(map[string](map[string]interface{}))
	testDataMap["metricOne"] = testData1
	testDataMap["metricTwo"] = testData2
	ret := RepoInfo{
		RepoName: "opendigger",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	rets = append(rets, ret)

	ret2 := RepoInfo{
		RepoName: "opendigger2",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	rets = append(rets, ret2)

	ret3 := RepoInfo{
		RepoName: "opendigger3",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	rets = append(rets, ret3)

	err := downloadService.SetData(rets, "metricOne")
	if err != nil {
		return
	}
	err2 := downloadService.Download("./")
	if err2 != nil {
		return
	}
}

func TestCompareDownloadService(t *testing.T) {
	fmt.Println("TestCompareDownloadService：")
	downloadService := &CompareDownloadService{}
	testDates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testData2 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testDataMap := make(map[string](map[string]interface{}))
	testDataMap["metricOne"] = testData1
	testDataMap["metricTwo"] = testData2
	ret1 := RepoInfo{
		RepoName: "opendigger1",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	ret2 := RepoInfo{
		RepoName: "opendigger2",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	err := downloadService.SetData(ret1, ret2, "html_output_compare")
	if err != nil {
		return
	}
	err2 := downloadService.Download()
	if err2 != nil {
		return
	}
}
