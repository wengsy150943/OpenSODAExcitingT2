/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"bufio"
	"exciting-opendigger/service"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strings"
)

var inputFile string
var outputFile string
var filter string
var sortKey string
var asc bool
var desc bool

var spokenLanguage string  //e.g. chinese,english
var programLanguage string //e.g. c++,java
var dateRange string       //daily、weekly、Monthly

// versionCmd represents the version command
var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Select a batch of data from OpenDigger",
	Long:  "Select a batch of data from OpenDigger",
	Run: func(cmd *cobra.Command, args []string) {
		//看一下是检索top的还是从文件检索
		ifTop := inputFile == "TOP"

		var repoList []string
		if ifTop {
			// TODO qk: call crawler
			var crawlerService service.CrawlerService
			crawlerService = &service.CrawlTrendingService{}
			//options用于根据前端接受的参数设置爬虫选项，分别是日期、
			opts := make([]service.Option, 0)

			if dateRange == "daily" {
				opts = append(opts, service.WithDaily())

			} else if dateRange == "weekly" {
				opts = append(opts, service.WithWeekly())

			} else if dateRange == "monthly" {
				opts = append(opts, service.WithMonthly())
			} else {
				log.Fatalf(" unvalid dateRange,please choose dateRange in daily or weekly or monthly")
			}

			opts = append(opts, service.WithProgramLanguage(programLanguage))

			_, err1 := service.SpokenLangCode[spokenLanguage]

			if !err1 {
				log.Fatalf(" unvalid spokenLanguage,please check the optional spokenLanguage in github.com")
			}
			opts = append(opts, service.WithSpokenLanguage(spokenLanguage))
			crawlerService.LoadOptions(opts...)
			res, err2 := crawlerService.Crawl()

			if err2 != nil {
				log.Fatalf(" crawl top failed,please check your Internet or check your input")
			}

			for _, x := range res {
				repoList = append(repoList, x.Link)
			}
		} else {
			if dateRange != "" || programLanguage != "" || spokenLanguage != "" {
				log.Println("Parameters of dateRange/programLanguage/spokenLanguage will be ignore here.")
			}

			file, e := os.Open(inputFile)
			if e != nil {
				log.Fatalf(inputFile + " file not found!")
			}
			defer file.Close()
			buffer := bufio.NewReader(file)
			for {
				repo, _, e := buffer.ReadLine()
				if e == io.EOF {
					break
				}
				repoList = append(repoList, string(repo))
			}
		}

		// 如果文件名本身已经带了csv后缀，把它裁掉
		outputFile = strings.TrimSuffix(outputFile, ".csv")

		service.DownLoadRepoList(repoList, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)

	batchCmd.Flags().StringVarP(&inputFile, "source", "s", "", "Repo source file")
	batchCmd.Flags().StringVarP(&outputFile, "position", "p", "", "Repo output file")

	batchCmd.Flags().StringVarP(&dateRange, "dateRange", "d", "", "TOP's dateRange")
	batchCmd.Flags().StringVarP(&programLanguage, "pl", "", "", "TOP's program language")
	batchCmd.Flags().StringVarP(&spokenLanguage, "sl", "", "", "TOP's spoken language")

	batchCmd.MarkFlagFilename("position")
	batchCmd.MarkFlagRequired("source")
	batchCmd.MarkFlagRequired("position")
}
