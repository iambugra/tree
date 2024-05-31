# tree

Simple Go implementation of `tree` command. Lists files and directories recursively in a
pretty format.

## usage

`go run main.go <path>`

Only `-a` flag allowed is for now, which lists all files and directories, including hidden
ones.

## examples

```shell
go run main.go -a sample_dir
sample_dir
├──sphinx
├──of
│  └──black
└──quartz
   └──judge
      └──.my
         └──vow

6 directories, 2 files
```

```shell
go run main.go sample_dir
sample_dir
├──how
│  └──vexingly
├──quick
└──daft
   ├──zebras
   └──jump

4 directories, 3 files
```
