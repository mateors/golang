## Go Tooling

## Topics to be covered
* Debugging
* Testing
* Profiling
* Tracing
* Unit Test
* [Go-wrk](https://github.com/tsliwowicz/go-wrk)
* [Go-torch](https://github.com/uber/go-torch)
* Benchmark

## Go formatting
> gofmt main.go \
> gofmt -w main.go \

## Cross compilation
> go build \ 
> go build -o name \
> GOOS=linux go build \
> GOOS=windows go build \
> file name \

## Compile & Install together ($GOPATH/bin)
> go install \
> $GOPATH/bin/name

## Downloading, Compiling & Installing into $GOPATH/bin
> go get github.com/mateors/reponame \
> go get -u https://github.com/golang/example/hello

## List all of your packages (dependencies)
> go list \
> go list -f '{{ .Name }}' \
> go list -f '{{ .Doc }}' \
> go list -f '{{ .Imports }}' \
> go list -f '{{ .Imports }}' strings \
> go list -f '{{ join .Imports "\n" }}' strings \
> go list -f '{{ join .Imports "\n" }}' strings \

## Go Documentation
> go doc strings \
> go doc strings Compare \

## Viewing documentation in the browser
> go doc -http := 8080



## Reference
* [Go Tooling in Action](https://www.youtube.com/watch?v=uBjoTxosSys)
* https://github.com/tsliwowicz/go-wrk
* https://github.com/uber/go-torch
