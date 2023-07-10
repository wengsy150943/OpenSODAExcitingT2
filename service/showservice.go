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

func PrintUserInfo(source_ UserInfo) {
	fmt.Printf("user.name: %s\n", source_.Username)

	for k,v := range source_.Data {
		jsonData,err := json.Marshal(v)
		if err != nil {
			fmt.Println("trans fail: ",err)
		}
		fmt.Printf("%s:%s\n", k, string(jsonData))
	}
}
