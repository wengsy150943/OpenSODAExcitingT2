package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
	// 模块和对应方法名暂时随便写一个，到时候替换具体实现的名字
	"opendigger/request"
	"opendigger/output"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "repo",
				Value: "",
				Usage: "request repository",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "month",
				Value: "",
				Usage: "request month",
			},
			&cli.StringFlag{
				Name:  "metric",
				Value: "",
				Usage: "request metrc",
			},
			&cli.BoolFlag{
				Name: "download, d",
				Value: false,
				Usage: "download",
			},
		},
		Action: func(c *cli.Context) error {
			repo := c.String("repo")
			month := c.String("month")
			metric := c.String("metric")
			download := c.Bool("download")

			var result map[string]string

			// 安全性检查

			// 至少要有一个参数
			if month == "" && metric == "" {
				fmt.Printf("Missing argue")
				return nil
			}

			// 获取数据指标
			// 模块和对应方法名暂时随便写一个，到时候替换具体实现的名字
			// 特定指标
			if month == "" {
				result = request.GetMetric(repo, metric)
			}
			// 特定月份的全部报告
			if metric == ""{
				result = request.GetMonth(repo, month)
			}
			// 特定月份的特定指标
			if month != "" && metric != "" {
				result = request.GetMonthAndMetric(repo, month, metric)
			}

			// 输出结果
			// 模块和对应方法名暂时随便写一个，到时候替换具体实现的名字
			// 直接输出结果

			// 导出文件
			if download {
				output.Download(result)
			} else { // 打印到屏幕
				output.Print(result)
			}


			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
