package service

import "testing"

func TestChart_Plot(t *testing.T) {
	a := Chart{}
	b := RepoInfo{}
	b.Getrepoinfo("X-lab2017/open-digger", "openrank")
	a.Plot("openrank", "bar", b.data)
	println(a.Title, a.Method, a.monthitems, a.numsitems)
	for _, c := range a.monthitems {
		println(c)
	}
	for _, d := range a.numsitems {
		println(d)
	}
}
