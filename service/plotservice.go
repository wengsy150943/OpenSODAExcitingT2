package service

import "encoding/json"

type PlotChartInterface interface {
	Plot(metric string, plotMethod string, body []byte)
}
type Chart struct {
	Method     string
	Title      string
	monthitems []string
	numsitems  []float64
}

func (c *Chart) Plot(metric string, plotMethod string, body []byte) {
	c.Method = plotMethod
	c.Title = metric + " " + plotMethod + "图"
	var v interface{}
	json.Unmarshal(body, &v)
	data := v.(map[string]interface{})
	for k, value := range data {
		c.monthitems = append(c.monthitems, k)
		c.numsitems = append(c.numsitems, value.(float64))
	}
}
