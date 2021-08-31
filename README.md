# golang
Golang Learning resources

## Data type detection
dataFields interface{}
> dtype := reflect.TypeOf(dataFields).Kind().String()

var c Company
> fmt.Println(reflect.ValueOf(&c).Kind().String()) //ptr \
> fmt.Println(reflect.ValueOf([]string{"aid", "cid", "type"}).Kind().String()) //slice

## Linux terminal build for windows
> `Syntax: env GOOS=target-OS GOARCH=target-architecture go build package-import-path`\
> $ env GOOS=windows go build

## Windows terminal build
> set GOOS=linux\
> go build -o appName

## Golang How to upgrade to Go 1.17 
> go mod tidy -go=1.17

## Module on by default
> go env -w GO111MODULE=auto

## Go Module Create
> `mkdir encdec && cd encdec` \
> `go mod init mateors.com/encdec` \
> write your source code & save the file

## Go Module, Call your code from anothe module
```go
import (
     ...
    "mateors.com/encdec"
)
```
> go mod edit -replace mateors.com/encdec=../encdec \
> go get mateors.com/encdec

## Resource
* https://golang.org/doc/tutorial/call-module-code
* https://go.dev/blog/go116-module-changes
* https://maelvls.dev/go111module-everywhere/
