# Golang Interface

## Hexagonal Architecture
Make sure we devide our software in such a way that each peice of the software maintain its separation of concerns
by doing this make it extreamly modular. 

1. App and domain logic in the middle
2. Ports & Adapters
3. User interface
4. Repo
5. External API
6. Message Queue
7. REST | GraphQL

## Decoupling

```go
package main

import (
	"fmt"
)

type reader interface {
	read() (int, error)
}

type file struct {
	name        string
	fileContent string
}

func (f *file) read() (int, error) 
	content := "file contents"
	f.fileContent = content
	return len(content), nil
}

type pipe struct {
	name        string
	pipeMessage string
}

func (p *pipe) read() (int, error) {
	msg := `{name: "Henry", title: "developer"}`
	p.pipeMessage = msg
	return len(msg), nil
}

func main() {
	f := file{
		name: "data.json",
	}
	p := pipe{
		name: "pipe_message",
	}

	retrieve(&f)
	fmt.Println(f.fileContent)
	retrieve(&p)
	fmt.Println(p.pipeMessage)
}

func retrieve(r reader) error {
	len, _ := r.read()
	fmt.Println(fmt.Sprintf("read %d bytes", len))
	return nil
}
```


### Package used in this project
* github.com/teris-io/shortid
* github.com/vmihailenco/msgpack
* gopkg.in/dealancer/validate.v2

## Resource
* [Interface](https://www.youtube.com/watch?v=qJKQZKGZgf0)
* [Hexagonal Microservices with Go](https://www.youtube.com/watch?v=rQnTtQZGpg8)
* [MessagePack](https://msgpack.org/index.html)
* [Golang Tensor Programming](https://www.youtube.com/watch?v=QyBXz9SpPqE&list=PLJbE2Yu2zumCe9cO3SIyragJ8pLmVv0z9)
* [understand-go-golang-interfaces](https://duncanleung.com/understand-go-golang-interfaces)
