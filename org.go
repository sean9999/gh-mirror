package main

import (
	"strings"
	"fmt"
	gh "github.com/cli/go-gh/v2"
)

type Org struct {
	Name string
}

// gh repo list devops-rona --json name,url,id

func (o Org) Repos() []Repo {
	stdOut, _, err := gh.Exec("repo", "list", o.Name, "--json", "name,url,id")
	if err != nil {
		return nil
	}

	repos := make([]Repo, 0)
	return repos
}


func firstPart(s string) string {
	pieces := strings.Split(s, " ")
	if len(pieces) > 0 {
		return pieces[0]
	}
	return ""
}

func OrgsFromString(s string) []Org {
	names := strings.Split(s, "\n")
	orgs := make([]Org, 0, len(names))
	for _, name := range names {

		if len(name) > 0 {
			orgs = append(orgs, Org{name})
		}
	}
	return orgs
}

