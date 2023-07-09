package service

import "testing"

func TestDownLoadRepoList(t *testing.T) {
	repoList := []string{"X-lab2017/open-digger", "X-lab2017/open-research", "X-lab2017/open-wonderland", "X-lab2017/open-leaderboard"}
	DownLoadRepoList(repoList, "test")
}
