# Mirror

This is a github cli (gh) plugin for mirroring all the repos
in an organization of set of organizations.

It will create a folder for every organization you're authenticated against,
and then clone or sync all the repos in the org folder.

If repos don't exist locally, they'll be cloned. If they do, they'll be synced.

synced means local tips will be updated to remote tips. Like `git pull`.

In a given repo, if it's in a dirty state
or contains commits that have not been pushed, it will refuse to act.

Install it like this:

```sh
$ gh ext install github.com/sean9999/gh-mirror
```

Use it like this:

```sh
$ gh mirror /path/to/mirror
```

The mirror will be created if you point to a folder that doesn't yet exist.
