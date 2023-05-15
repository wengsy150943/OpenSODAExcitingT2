package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"log"
	"os"
)

func PlotBar(metric string, body []byte) {
	monthItems := []string{}
	numsItems := []float64{}

	var v2 interface{}

	json.Unmarshal(body, &v2)
	data2 := v2.(map[string]interface{})
	for k, v := range data2 {
		fmt.Println(k)
		monthItems = append(monthItems, k)
		numsItems = append(numsItems, v.(float64))
	}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(monthItems).
		AddYAxis(metric+"值", numsItems)
	f, err := os.Create("../assets/images/bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}
