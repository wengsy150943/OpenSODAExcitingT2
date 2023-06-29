package utils

import (
	"exciting-opendigger/service"
	"testing"
)

func TestCreate(t *testing.T) {
	Create("testdb")
}
func TestInsert(t *testing.T) {
	a := service.RepoInfo{}
	a.Getrepoinfo("X-lab2017/open-digger", "openrank", "")
	Insert("openrank")

}
