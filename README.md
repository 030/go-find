# go-find
A go equivalent to find

## library

In order to use go-find as a library, one could issue:

```
go get github.com/030/go-find
```

## tool

### usage

```
[user@localhost go-find]$ ./go-find -h
2018/12/12 02:37:06 Usage: go-find <dir e.g. ~/.gradle> <filename e.g. commons-io-1.2.3>
exit status 1
```

### example

```
[user@localhost go-find]$ go run main.go ~ hello.go
/home/user/hello.go
```