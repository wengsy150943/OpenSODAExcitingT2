package service

import (
	"fmt"
	"encoding/json"
)

var printList = []string{
	"repo.name",
	"repo.url",
}

func PrintRepoInfo(source_ RepoInfo) {
	fmt.Printf("%s: %s\n", printList[0], source_.repoName)
	fmt.Printf("%s: %s\n", printList[1], source_.repoUrl)

	for _,v := range Metrics {
		datum, ok := source_.data[v]
		if ok{
			fmt.Printf("%s:", v)
			jsonData,err := json.Marshal(datum)
			if err != nil {
				fmt.Errorf("trans fail",err)
			}
			fmt.Println(string(jsonData))
		}
	}

}