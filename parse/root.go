package parse

import (
	"exciting-opendigger/service"
	"exciting-opendigger/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// 访问参数
type Query struct {
	repo, user, month, metric string
}



var queryPara Query
var repoInfo service.RepoInfo
var repoInfoCompare service.RepoInfo

var userInfo service.UserInfo
var userInfoCompare service.UserInfo
var isShow bool = true

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exciting-opendigger",
	Short: "Exciting-opendigger,a tool to get and download opendigger result",
	Long:  `Exciting-opendigger,a tool to get and download opendigger result.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	utils.Insertlog(strings.Join(os.Args[1:]," "))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 更新查询参数，无论是显示还是下载都需要这些参数
	rootCmd.PersistentFlags().StringVarP(&queryPara.repo, "repo", "r", "", "Repository asked")
	rootCmd.PersistentFlags().StringVarP(&queryPara.user, "user", "u", "", "User asked")

	rootCmd.MarkFlagsMutuallyExclusive("repo", "user")

	rootCmd.PersistentFlags().StringVarP(&queryPara.month, "month", "M", "", "Month of asked metric")
	rootCmd.PersistentFlags().StringVarP(&queryPara.metric, "metric", "m", "", "Metric asked")

	// 因为user不支持任何参数，在传参的时候预先卡死
	rootCmd.MarkFlagsMutuallyExclusive("user", "month")
	rootCmd.MarkFlagsMutuallyExclusive("user", "metric")
}
