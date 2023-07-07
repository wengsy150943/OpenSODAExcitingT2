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
		service.DownLoadRepoList(repoList, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)

	batchCmd.Flags().StringVarP(&inputFile, "source", "s", "", "Repo source file")
	batchCmd.Flags().StringVarP(&outputFile, "position", "p", "", "Repo output file")
	batchCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter for metric")
	batchCmd.Flags().StringVarP(&sortKey, "orderBy", "o", "", "Sort key")
	batchCmd.Flags().BoolVar(&asc, "asc", true, "Sort ASC")
	batchCmd.Flags().BoolVar(&desc, "desc", false, "Sort DESC")

	batchCmd.MarkFlagFilename("position")
	batchCmd.MarkFlagRequired("source")
	batchCmd.MarkFlagRequired("position")
	batchCmd.MarkFlagsMutuallyExclusive("asc", "desc")
}
