package service

import (
	"encoding/json"
	"fmt"
)

var printList = []string{
	"repo.name",
	"repo.url",
}

func PrintRepoInfo(source_ RepoInfo) {
	fmt.Printf("%s: %s\n", printList[0], source_.RepoName)
	fmt.Printf("%s: %s\n", printList[1], source_.RepoUrl)

	for _,v := range Metrics {
		datum, ok := source_.Data[v]
		if ok{
			jsonData,err := json.Marshal(datum)
			if err != nil {
				fmt.Println("trans fail: ",err)
			}
			fmt.Printf("%s:%s\n", v, string(jsonData))
		}
	}
}
