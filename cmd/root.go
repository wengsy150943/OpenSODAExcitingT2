package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// 访问参数
type Query struct {
	repo,month,metric string
}
var queryPara Query

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exciting-opendigger",
	Short: "exciting-opendigger,a tool to get and download opendigger result",
	Long:  `exciting-opendigger,a tool to get and download opendigger result.`,
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
	rootCmd.MarkFlagRequired("repo")
	rootCmd.PersistentFlags().StringVarP(&queryPara.month, "month", "M", "", "Month of asked metric")
	rootCmd.PersistentFlags().StringVarP(&queryPara.metric, "metric", "m", "", "Metric asked")
}