package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/cli/go-gh/v2"
)

type Org struct {
	Name string
}

func (o Org) Repos() ([]Repo, error) {
	stdOut, _, err := gh.Exec("repo", "list", o.Name, "--json", "name,url,sshUrl,id")
	if err != nil {
		return nil, err
	}
	repos := make([]Repo, 0)
	err = json.Unmarshal(stdOut.Bytes(), &repos)
	return repos, err
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

func GetOrgs() []Org {
	args := []string{"org", "list"}
	stdOut, _, err := gh.Exec(args...)
	if err != nil {
		log.Fatal(err)
	}
	orgs := OrgsFromString(stdOut.String())
	return orgs
}
