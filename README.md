# git-purge

git-purge is git sub command to delete merged branches.

THIS APPLICATION IS VERY EXPERIMENTAL.
THERE ARE NO GUARANTEE FOR RESULT OF IT.
YOU SHOULDN'T USE IT IF YOU CANNOT UNDERSTAND ITS BEHAVIOR.

## Installation

```
go get github.com/handlename/git-purge
```

## Usage

To delete all merged branches, run `git purge` without any options.

```
$ git purge
```

With `--dryrun` option, list branches and do nothing more.

```
$ git purge --dryrun
```

If you want to delete remote branches too, try `--remote` option.

```
$ git purge --remote
```

`--ignore "<regexp>"` ignores branches matches `<regexp>`.

```
$ git purge --ignore "keep/.*"
```

## Licence

MIT

## Author

[handlename](https://github.com/handlename)
