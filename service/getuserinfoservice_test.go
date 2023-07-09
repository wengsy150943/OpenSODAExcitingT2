package service

import "testing"

func TestGetCertainUser(t *testing.T) {
	testData := make(map[string]interface{})

	username := "frank-zsy"
	a := GetCertainUser(username)
	println(a.Data["openrank"]["2015-02"].(float64))
	println(a.Data["developernetwork"]["node"].([]interface{})[0].([]interface{})[0].(string))
	if a.Data["openrank"]["2015-02"].(float64) != 0.64 || a.Data["developernetwork"]["node"].([]interface{})[0].([]interface{})[0].(string) != "snyk-bot" || a.Data["developernetwork"]["node"].([]interface{})[0].([]interface{})[0].(float64) != 10833.52 {

	}
}
