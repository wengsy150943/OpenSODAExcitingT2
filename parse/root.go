package parse

import (
	"exciting-opendigger/service"
	"os"
	"github.com/spf13/cobra"
)

// 访问参数
type Query struct {
	repo, month, metric string
}



var queryPara Query
var repoInfo service.RepoInfo
var repoInfoCompare service.RepoInfo

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exciting-opendigger",
	Short: "Exciting-opendigger,a tool to get and download opendigger result",
	Long:  `Exciting-opendigger,a tool to get and download opendigger result.`,
	PersistentPostRun: func (cmd *cobra.Command, args []string)  {
		
	}, 
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 更新查询参数，无论是显示还是下载都需要这些参数
	rootCmd.PersistentFlags().StringVarP(&queryPara.repo, "repo", "r", "", "Repository asked")
	rootCmd.PersistentFlags().StringVarP(&queryPara.month, "month", "M", "", "Month of asked metric")
	rootCmd.PersistentFlags().StringVarP(&queryPara.metric, "metric", "m", "", "Metric asked")
}
