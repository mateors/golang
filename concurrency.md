# Golang Concurrency

```go
package main

import (
	"fmt"
)

func main() {

	c1 := make(chan string)
  	c1<- "Mostain" //sending
  
  	//code blocks here until something is ready to receive
	
	v := <-c1 //receiving
	fmt.Println(v)

}
```
* The above won't run because of deadlock
* Because of both are in the same goroutine

You can fix the above problem in two ways

### Way-01, wrap the sending code in a separate goroutine
```go
	go func() {
		c1 <- "mostain"
	}()

	v := <-c1
	fmt.Println(v)
```

### Way-02, using the buffered channel
```go
	c1 := make(chan string,1) //making a buffered channel by giving a capacity
  	c1<- "Mostain" //sending
  
	v := <-c1 //receiving
	fmt.Println(v)
	
```
### Channel select
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			c1 <- "worker1 every 500 ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {

		for i := 0; i < 3; i++ {
			c2 <- "worker2 every 1.5 seconds"
			time.Sleep(time.Millisecond * 1500)
		}

	}()

	var counter int
	for {

		//fmt.Println(<-c1)
		//fmt.Println(<-c2)

		select {
		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)
			counter++
		default:
			fmt.Println("none of them are ready to send")
			time.Sleep(time.Millisecond * 100)
		}

		//fmt.Println("counter:", counter)
		if counter == 3 {
			break
		}
	}

}
```

## Worker pool

```go
package main

import (
	"fmt"
	"sync"
)

func fib(n int) int {

	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func worker(jobs <-chan int, result chan<- int) {

	for n := range jobs {
		result <- fib(n)
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 100; i++ {
		res := <-results
		fmt.Println(res)
	}

}
```


## Resource
* [Get-a-taste-of-concurrency](https://levelup.gitconnected.com/get-a-taste-of-concurrency-in-go-625e4301810f)
* [Multihreading-and-multiprocessing](https://www.mineiros.io/blog/guide-to-multihreading-and-multiprocessing)
* [Common-mistakes or gotchas](https://www.linkedin.com/pulse/50-shades-go-traps-gotchas-common-mistakes-new-golang-okhotnikov-1f/)
