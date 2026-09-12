package main

import (
	"fmt"
	"log"

	gh "github.com/cli/go-gh/v2"
)

func main() {
	args := []string{"org", "list"}
	stdOut, stdErr, err := gh.Exec(args...)
	if err != nil {
		fmt.Println(stdErr.String())
		log.Fatal(err)
	}


	orgs := OrgsFromString(stdOut.String())

	fmt.Println(orgs)


	repos := orgs[0].Repos()

	for _, repo := range repos {

		fmt.Println(repo.Name)

	}

	//a,_,_ := gh.Exec("repo", "list", "devops-rona")

	//fmt.Println(a.String())


}
