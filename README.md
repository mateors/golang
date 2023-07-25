# golang
## Golang Learning resources
> Software engineering is a craft that should be honed, debated, and continuously improved.

## Golang Installation or Upgradation
> go version \
> which go


> wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz

> sudo rm -rf /usr/local/go

> sudo tar -xzf go1.17.7.linux-amd64.tar.gz

> sudo mv go /usr/local

> export PATH=$PATH:/usr/local/go/bin

> source $HOME/.profile

> go version


## Data type detection
dataFields interface{}
> dtype := reflect.TypeOf(dataFields).Kind().String()

var c Company
> fmt.Println(reflect.ValueOf(&c).Kind().String()) //ptr \
> fmt.Println(reflect.ValueOf([]string{"aid", "cid", "type"}).Kind().String()) //slice

## How to get the name of a struct?
```go
func structName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

type Achead struct{
 ID int
 Code string
}

fmt.Println(structName(Achead{}))
fmt.Println(structName(&Achead{})) //Achead
```

## Linux terminal build for windows
> `Syntax: env GOOS=target-OS GOARCH=target-architecture go build package-import-path`\
> $ env GOOS=windows go build

## Windows terminal build
> set GOOS=linux\
> go build -o appName

## Windows powershell build
If you're using PowerShell, then changing values of environment variables should be done like: `$Env:<variable-name> = "<new-value>"`

> $Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build -o hello

### Cache clean command
> go clean --cache

### Trouble-shooting
Cleaning the go caches once fixed some “unknown import path”/“cannot find module providing package” errors
> go clean -cache -modcache -i -r

### Testing command
> $Env:GOOS="windows";go test .\utility\

> `$Env:GOOS="windows";go test -timeout 30s -run ^TestSsPLNT$ graphmysql/utility`

> $Env:GOOS="windows";go test ./...

## Golang How to upgrade to Go 1.17 
> go mod tidy -go=1.17
> go mod edit -go 1.20

## Module on by default
> go env -w GO111MODULE=auto

> fmt.Println("NumGoroutine:", runtime.NumGoroutine()) \
> GODEBUG=gctrace=1 ./chaldal

## Go Module Create
> `mkdir encdec && cd encdec` \
> `go mod init mateors.com/encdec` \
> write your source code & save the file

### This command deletes the cache downloaded along with unpacked code dependencies
> go clean -modcache

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
* https://hasura.io/graphql/database/
* https://nats.io/
* https://github.com/hasura/graphql-engine
* https://hasura.io/graphql/database/
* https://lecstor.com/go-cheatsheet
