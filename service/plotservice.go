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
	f, err := os.Create("../assets/images/" + metric + "_bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}

func Plotline(metric string, body []byte) {
	monthItem := []string{}
	numsItem := []float64{}

	var v interface{}
	json.Unmarshal(body, &v)
	data := v.(map[string]interface{})

	for k, v := range data {
		monthItem = append(monthItem, k)
		numsItem = append(numsItem, v.(float64))
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{
		Title: metric + "折线图",
	}, charts.ToolboxOpts{Show: true})
	line.AddXAxis(monthItem).AddYAxis(metric+"值", numsItem)
	f, err := os.Create("../assets/images/" + metric + "_line.html")
	if err != nil {
		log.Println(err)
	}
	line.Render(f)
}
