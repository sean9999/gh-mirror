package main

import (
	"fmt"
	"os"
	"strings"
)

const RootDir = "/Users/6012777/repos/rona/mirroir"

func main() {
	orgs := GetOrgs()
	for _, org := range orgs {
		fmt.Println(org.Name)
		repos, err := org.Repos()
		if err != nil {
			panic(err)
		}
		for _, repo := range repos {
			fmt.Println(repo.Name)
			myDir := strings.Join([]string{RootDir, org.Name, repo.Name}, string(os.PathSeparator))
			err := EnsureSynced(org, repo, myDir)
			fmt.Println(err)
		}
	}
}
