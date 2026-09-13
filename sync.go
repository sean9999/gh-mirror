package main

import (
	"errors"

	"github.com/armed/mkdirp"
	"github.com/nodefortytwo/isgit"
	"github.com/sean9999/hermeti"
)

func sync(env *hermeti.Env, org Org, repo Repo) error {
	return runCli(env.OutStream, env.ErrStream, "repo", "sync")
}

func clone(env *hermeti.Env, repo Repo) error {
	return runCli(env.OutStream, env.ErrStream, "repo", "clone", repo.SshUrl, ".")
}

// EnsureSynced ensures a folder is a git repo and is synced to upstream,
// cloning if necessary.
func EnsureSynced(env *hermeti.Env, org Org, repo Repo, dir string) error {
	err := EnsureDir(env, dir)
	if err != nil {
		return err
	}
	isRepo, err := isgit.IsGitRepo(dir)
	if err != nil {
		return err
	}
	if isRepo {
		return sync(env, org, repo)
	}
	return clone(env, repo)
}

// EnsureDir ensures a directory exists by creating it or making sure it's already there.
// It also goes (chdir) into it.
func EnsureDir(env *hermeti.Env, dir string) error {
	info, err := env.Filesystem.Stat(dir)
	if err != nil {
		err = mkdirp.Mk(dir, 0755)
		if err != nil {
			return err
		}
		return env.Chdir(dir)
	}
	if info.IsDir() == false {
		return errors.New("not a dir")
	}
	return env.Chdir(dir)
}
