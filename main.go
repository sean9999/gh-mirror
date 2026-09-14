package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sean9999/hermeti"
)

var _ hermeti.Runner = (*state)(nil)

type state struct {
	rootDir string // rootDir is where the mirror should live
}

func (a *state) Run(env *hermeti.Env) {
	if len(env.Args) < 2 {
		a.rootDir = "."
	} else {
		a.rootDir = env.Args[1]
	}
	err := EnsureDir(env, a.rootDir)
	if err != nil {
		fmt.Fprintln(env.ErrStream, "You must pass in a valid directory")
		panic(err)
	}

	orgs := GetOrgs()

	fmt.Fprintf(env.OutStream, "syncing %s...\n", a.rootDir)
	for _, org := range orgs {
		repos, err := org.Repos(env)
		if err != nil {
			panic(err)
		}
		for _, repo := range repos {
			fmt.Fprintf(env.OutStream, "%s\t%s\n", repo.Org.Name, repo.Name)
			myDir := strings.Join([]string{a.rootDir, org.Name, repo.Name}, string(os.PathSeparator))
			err := EnsureSynced(env, repo, myDir)
			if err != nil {
				fmt.Fprintln(env.ErrStream, err)
			}
		}
	}
}

func main() {
	app := new(state)
	cli := hermeti.NewRealCli(app)
	cli.Run()
}
