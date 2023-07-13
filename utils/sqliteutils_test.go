package utils

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

func TestInsertAndRead(t *testing.T) {
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3.0, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testQueryresult := make(map[string](map[string]interface{}))
	metric := "activity"
	testQueryresult[metric] = testData1

	res := CachedRepoInfo{}
	err := ReadQuerySingleMetric(&res, "X-lab2017/open-digger", metric)
	//如果在缓存中没有查询到，会返回Errrecordnotfound错误，但会导致test报错，if 里的内容也会执行
	//但test会返回报错。
	if errors.Is(err, gorm.ErrRecordNotFound) {
		dates := make([]string, len(testData1))
		cnt := 0
		for k, _ := range testData1 {
			dates[cnt] = k
			cnt++
		}
		err1 := InsertSingleQuery("X-lab2017/open-digger", "https://oss.x-lab.info/open_digger/github/X-lab2017/open-digger", "activity", "", dates, testQueryresult)
		if err1 != nil {
			t.Fatal(err)
		}
	}
}

func TestInsertsinglequery(t *testing.T) {
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3.0, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}

	testQueryresult := make(map[string](map[string]interface{}))

	testQueryresult["openrank"] = testData1

	dates := make([]string, len(testData1))
	cnt := 0
	for k, _ := range testData1 {
		dates[cnt] = k
		cnt++
	}

	err := InsertSingleQuery("X-lab2017/open-digger", "https://oss.x-lab.info/open_digger/github/X-lab2017/open-digger", "openrank", "", dates, testQueryresult)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadquery(t *testing.T) {
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3.0, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testQueryresult := make(map[string](map[string]interface{}))
	metric := "openrank"
	testQueryresult[metric] = testData1

	res := CachedRepoInfo{}
	ReadQuerySingleMetric(&res, "X-lab2017/open-digger", metric)
	println(res.Data["openrank"]["2023-06"])
	for k, v := range testData1 {
		if res.Data[metric][k].(float64) != v {
			t.Errorf("Read query error" + k + strconv.FormatFloat(float64(res.Data[metric][k].(float64)), 'f', 6, 64) + " correct:" + strconv.FormatFloat(v.(float64), 'f', 6, 64))
		}
	}
}

func TestUpdateSingleRow(t *testing.T) {
	//添加2023-06:17.5项
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3.0, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2023-06": 17.5, "2021-10-raw": 2.84}
	testQueryresult := make(map[string](map[string]interface{}))
	metric := "openrank"
	testQueryresult[metric] = testData1
	dates := make([]string, len(testData1))
	cnt := 0
	for k, _ := range testData1 {
		dates[cnt] = k
		cnt++
	}

	err := UpdateSingleRow("X-lab2017/open-digger", "openrank", dates, testQueryresult)
	if err != nil {
		t.Error("UPDATE FAILED")
	}
}
func TestInsertlog(t *testing.T) {
	testlog1 := "opendigger repo = X-lab2017/open-digger metric = OpenRank month = 2023-01"
	testlog2 := "opendigger repo = X-lab2017/open-digger metric = OpenRank"
	testlog3 := "opendigger repo = X-lab2017/open-digger month = 2023-01"
	err := Insertlog(testlog1)
	err = Insertlog(testlog2)
	err = Insertlog(testlog3)
	if err != nil {
		t.Fatal("Insert log error")
	}
}
func TestReadlog(t *testing.T) {
	logs := []Searchhistory{}
	Readlog(&logs)
	testlogs := []string{
		"opendigger repo = X-lab2017/open-digger metric = OpenRank month = 2023-01",
		"opendigger repo = X-lab2017/open-digger metric = OpenRank",
		"opendigger repo = X-lab2017/open-digger month = 2023-01",
	}
	for i, log := range testlogs {
		if logs[i].Log != log {
			t.Errorf("Read logs error " + "\"" + logs[i].Log + "\"" + " Correct is :" + "\"" + log + "\"")
		}
	}
}

func TestInsertUserInfo(t *testing.T) {
	temp := CachedUserInfo{}
	CreateTable(temp)
	testData1 := map[string](interface{}){
		"2015-02":     0.64,
		"2015-03":     0.42,
		"2015-04":     0.36,
		"2015-05":     0.45,
		"2015-06":     0.38,
		"2015-07":     0.32,
		"2015-08":     0.28,
		"2015-09":     0.23,
		"2015-10":     0.2,
		"2015-11":     0.17,
		"2015-12":     0.14,
		"2016-01":     0.12,
		"2016-02":     0.1,
		"2016-03":     0.09,
		"2016-05":     0.61,
		"2016-06":     0.52,
		"2016-07":     0.44,
		"2016-08":     0.38,
		"2016-09":     0.32,
		"2016-10":     0.27,
		"2016-11":     0.23,
		"2016-12":     0.2,
		"2017-01":     0.17,
		"2017-02":     0.14,
		"2017-03":     0.12,
		"2017-04":     0.1,
		"2017-05":     0.09,
		"2018-02":     1.99,
		"2018-03":     1.43,
		"2018-04":     0.99,
		"2018-05":     0.67,
		"2018-06":     0.47,
		"2018-07":     0.8,
		"2018-08":     1.97,
		"2018-09":     1.68,
		"2018-10":     1.06,
		"2018-11":     1,
		"2018-12":     1.73,
		"2019-01":     3.02,
		"2019-02":     3.06,
		"2019-03":     2.9,
		"2019-04":     3.9,
		"2019-05":     2.14,
		"2019-06":     1.82,
		"2019-07":     1.02,
		"2019-08":     2.22,
		"2019-09":     2.26,
		"2019-10":     1.5,
		"2019-11":     1.6,
		"2019-12":     1.44,
		"2020-01":     10.41,
		"2020-02":     9.94,
		"2020-03":     8.57,
		"2020-04":     6.74,
		"2020-05":     5.34,
		"2020-06":     3.1,
		"2020-07":     2.76,
		"2020-08":     4.96,
		"2020-09":     5.91,
		"2020-10":     4.64,
		"2020-11":     3.49,
		"2020-12":     4.33,
		"2021-01":     5.54,
		"2021-02":     4.06,
		"2021-03":     3.51,
		"2021-04":     4.32,
		"2021-05":     4.11,
		"2021-06":     3.54,
		"2021-07":     2.27,
		"2021-08":     1.86,
		"2021-09":     1.58,
		"2021-10":     2,
		"2021-11":     2.61,
		"2021-12":     3.35,
		"2022-01":     4.94,
		"2022-02":     4.45,
		"2022-03":     5.68,
		"2022-04":     5.91,
		"2022-05":     4.85,
		"2022-06":     5.68,
		"2022-07":     5.2,
		"2022-08":     5.05,
		"2022-09":     6,
		"2022-10":     5.4,
		"2022-11":     7.54,
		"2022-12":     10.46,
		"2023-01":     12.48,
		"2023-02":     23.86,
		"2023-03":     16.71,
		"2023-04":     12.63,
		"2023-05":     9.47,
		"2023-06":     6.33,
		"2021-10-raw": 1.48,
	}
	dates := make([]string, len(testData1))
	for k, _ := range testData1 {
		dates = append(dates, k)
	}
	testQueryresult := make(map[string](map[string]interface{}))
	testQueryresult["openrank"] = testData1
	username := "frank-zsy"
	InsertUserInfo(username, testQueryresult, dates)

}

func TestReadSingleUserInfo(t *testing.T) {
	a := CachedUserInfo{}
	username := "frank-zsy"
	err := ReadSingleUserInfo(&a, username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		println("Not found")
	}
	println(a.Username)
	println(a.Data["openrank"]["2020-08"].(float64))
}
