/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"exciting-opendigger/service"
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
		// 看一下是检索top的还是从文件检索
		ifTop := inputFile == "TOP"
		
		if (ifTop){
			source = map[string]string{"sth":"sth"}
		}else{
			source = map[string]string{"sth":"sth"}
		}

		var downloadService service.BatchDownloadService
		str,_ :=json.Marshal(source)
		downloadService.SetData(string(str), outputFile)
		downloadService.Download()
	},
}





func init() {
	rootCmd.AddCommand(batchCmd)

	batchCmd.Flags().StringVarP(&inputFile, "source", "s", "", "Repo source file")
	batchCmd.Flags().StringVarP(&outputFile, "position", "p", "", "Repo output file")
	batchCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter for metric")
	batchCmd.Flags().StringVarP(&sortKey, "orderBy", "o", "", "Sort key")
	batchCmd.Flags().BoolVar(&asc,"asc",true,"Sort ASC")
	batchCmd.Flags().BoolVar(&desc,"desc",false,"Sort DESC")

	
	batchCmd.MarkFlagFilename("position")
	batchCmd.MarkFlagRequired("source")
	batchCmd.MarkFlagRequired("position")
	batchCmd.MarkFlagsMutuallyExclusive("asc","desc")

}
