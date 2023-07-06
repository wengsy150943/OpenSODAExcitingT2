package service

import (
	"strconv"
	"testing"
)

func TestGetCertainMetric(t *testing.T) {
	result := map[string]float32{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	a := GetRepoInfoOfMetric("X-lab2017/open-digger", "openrank")
	println(a.Data["openrank"]["2020-08"].(float64))
	for k, v := range result {
		if float32(a.Data["openrank"][k].(float64)) != v {
			temp := a.Data["openrank"][k].(float64)
			t.Errorf("get certain repo info false " + strconv.FormatFloat(float64(temp), 'f', 6, 64) + " -> " + strconv.FormatFloat(float64(v), 'f', 6, 64))
			break
		}
	}
}

func TestGetCertainRepo(t *testing.T) {
	result := map[string]float32{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}

	for k, v := range result {
		a := GetCertainRepoInfo("X-lab2017/open-digger", "openrank", k)
		if float32(a.Data["openrank"][k].(float64)) != v {
			temp := a.Data["openrank"][k].(float64)
			t.Errorf("get certain repo info false " + strconv.FormatFloat(float64(temp), 'f', 6, 64) + " -> " + strconv.FormatFloat(float64(v), 'f', 6, 64))
			break
		}
	}
}

func TestGetCertainRepoInfoNotInCache(t *testing.T) {
	result := map[string]float32{"2020-08": 10,
		"2020-09":     13,
		"2020-10":     10,
		"2020-11":     4,
		"2020-12":     4,
		"2021-01":     1,
		"2021-02":     7,
		"2021-03":     3,
		"2021-04":     6,
		"2021-05":     4,
		"2021-06":     5,
		"2021-07":     4,
		"2021-08":     1,
		"2021-09":     2,
		"2021-10":     4,
		"2021-11":     7,
		"2021-12":     3,
		"2022-01":     5,
		"2022-02":     5,
		"2022-03":     102,
		"2022-04":     3,
		"2022-05":     4,
		"2022-06":     1,
		"2022-07":     2,
		"2022-08":     7,
		"2022-09":     10,
		"2022-10":     9,
		"2022-11":     9,
		"2022-12":     3,
		"2023-01":     1,
		"2023-02":     8,
		"2023-03":     12,
		"2023-04":     6,
		"2023-05":     2,
		"2023-06":     3,
		"2021-10-raw": 1}
	a := GetRepoInfoOfMetric("X-lab2017/open-digger", "stars")
	println(a.Data["stars"]["2020-08"].(float64))
	for k, v := range result {
		if float32(a.Data["stars"][k].(float64)) != v {
			temp := a.Data["stars"][k].(float64)
			t.Errorf("get certain repo info false " + strconv.FormatFloat(float64(temp), 'f', 6, 64) + " -> " + strconv.FormatFloat(float64(v), 'f', 6, 64))
			break
		}
	}
}

func TestGetCertainRepoSpecial(t *testing.T) {
	avg := 149.29
	month := "2020-08"
	metric := "issue_response_time"
	a := GetCertainRepoInfo("X-lab2017/open-digger", metric, month)
	if a.Data["issue_response_time"][month].(map[string]interface{})["avg"] != avg {
		t.Errorf("get Certain Repo Special fail")
	}
}

func TestGetRepoInfoOfMonth(t *testing.T) {
	var result float32
	result = 4.5
	month := "2020-08"
	a := GetRepoInfoOfMonth("X-lab2017/open-digger", month)
	if float32(a.Data["openrank"][month].(float64)) != result {
		t.Errorf("get certain repoinfo of month false ")
	}

}
