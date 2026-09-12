package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/armed/mkdirp"
	"github.com/cli/go-gh/v2"
	"github.com/nodefortytwo/isgit"
)

func sync(org Org, repo Repo) error {
	args := []string{"repo", "sync"}
	stdOut, stdErr, err := gh.Exec(args...)

	fmt.Println(stdOut.String())
	fmt.Println(stdErr.String())

	return err
}

func clone(repo Repo) error {
	args := []string{"repo", "clone", repo.SshUrl, "."}
	stdOut, stdErr, err := gh.Exec(args...)
	fmt.Println(stdOut.String())
	fmt.Println(stdErr.String())
	return err
}

func EnsureSynced(org Org, repo Repo, dir string) error {
	err := EnsureDir(dir)
	if err != nil {
		return err
	}
	isRepo, err := isgit.IsGitRepo(dir)
	if err != nil {
		return err
	}
	if isRepo {
		return sync(org, repo)
	}
	return clone(repo)
}

func EnsureDir(dir string) error {

	info, err := os.Stat(dir)
	if err != nil {
		err = mkdirp.Mk(dir, 0755)
		if err != nil {
			return err
		}
		return os.Chdir(dir)
	}
	if info.IsDir() == false {
		return errors.New("not a dir")
	}
	return os.Chdir(dir)
}
