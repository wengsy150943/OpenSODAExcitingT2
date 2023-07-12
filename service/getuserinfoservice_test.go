package service

import (
	"exciting-opendigger/utils"
	"testing"
)

func TestGetCertainUser(t *testing.T) {
	username := "frank-zsy"
	a := GetCertainUser(username)
	//println(a.Data["openrank"]["2015-02"].(float64))
	//println(a.Data["developernetwork"]["nodes"].([]interface{})[0].([]interface{})[0].(string))
	if a.Data["openrank"]["2015-02"].(float64) != 0.64 || a.Data["developernetwork"]["nodes"].([]interface{})[0].([]interface{})[0].(string) != "snyk-bot" || a.Data["developernetwork"]["nodes"].([]interface{})[0].([]interface{})[1].(float64) != 10833.52 {
		t.Errorf("Get userinfo failed")
	}
	for _, k := range a.Dates {
		println(k)
	}
}

func TestParseuser(t *testing.T) {
	a := GetCertainUser("frank-zsy")
	b := utils.Usermetric{}
	b = utils.Parseuser(a.Data, b)
	println(b.Developernetwork["nodes"][0].([]interface{})[0].(string))
	if b.Developernetwork["nodes"][0].([]interface{})[0].(string) != "snyk-bot" {
		t.Errorf("parse error")
	}
}
