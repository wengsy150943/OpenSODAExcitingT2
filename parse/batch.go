/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	"bufio"
	"exciting-opendigger/service"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string
var filter string
var sortKey string
var asc bool
var desc bool

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
		} else {
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
		outputFile,_ = strings.CutSuffix(outputFile,".csv")
		service.DownLoadRepoList(repoList, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)

	batchCmd.Flags().StringVarP(&inputFile, "source", "s", "", "Repo source file")
	batchCmd.Flags().StringVarP(&outputFile, "position", "p", "", "Repo output file")

	batchCmd.MarkFlagFilename("position")
	batchCmd.MarkFlagRequired("source")
	batchCmd.MarkFlagRequired("position")
}
