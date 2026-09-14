# Mirror

This is a GitHub CLI (gh) plugin for mirroring all the repos
in an organization, or set of organizations.

It will create a folder for every organization you're authenticated against,
and then clone or sync all the repos in the org folder.

If repos don't exist locally, they'll be cloned. If they do, they'll be [synced](https://cli.github.com/manual/gh_repo_sync).

If a repo's in a dirty state or contains commits that have not been pushed,
it will refuse to act. It's non destructive.

Install it like this:

```sh
$ gh ext install github.com/sean9999/gh-mirror
```

Use it like this:

```sh
$ gh mirror /path/to/mirror
```

If the folder doesn't exist, it'll be created. If you don't pass in a folder argument, current directory is assumed.
