/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package parse

import (
	// "exciting-opendigger/service"
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
		// ifTop := inputFile == "TOP"
		
		// var repoList []string
		// if (ifTop){
		// 	// TODO qk: call crawler
		// }else{
		// 	// TODO syz lh: read file
		// }

		// // TODO syz lh: get repo info
		// repoInfoList := service.GetRepoList(repoList)


		// var downloadService service.DownloadService
		// downloadService = &service.BatchDownloadService{}
		// downloadService.SetData(repoInfoList, outputFile)
		// downloadService.Download()
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
