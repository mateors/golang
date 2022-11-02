# Golang Concurrency

```go
package main

import (
	"fmt"
)

func main() {

	c1 := make(chan string)
  c1<- "Mostain" //sending
  
	v := <-c1 //receiving
	fmt.Println(v)

}
```
* The above won't run because of deadlock
* Because of both are in the same goroutine



## Resource
* [Get-a-taste-of-concurrency](https://levelup.gitconnected.com/get-a-taste-of-concurrency-in-go-625e4301810f)
* [Multihreading-and-multiprocessing](https://www.mineiros.io/blog/guide-to-multihreading-and-multiprocessing)
* [Common-mistakes or gotchas](https://www.linkedin.com/pulse/50-shades-go-traps-gotchas-common-mistakes-new-golang-okhotnikov-1f/)
