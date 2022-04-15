# Golang Lazy Evaluation using closure
> What is a closure in computing? \
In computer science, a closure is a function that has an environment of its own. In this environment, there is at least one bound variable (a name that has a value, such as a number). The closure's environment keeps the bound variables in memory between uses of the closure.

## Higher-order function
In mathematics and computer science, a higher-order function is a function that does at least one of the following: takes one or more functions as arguments, returns a function as its result. All other functions are first-order functions. In mathematics higher-order functions are also termed operators or functionals.

A Higher-Order function is a function that receives a function as an argument or returns the function as output. Higher order functions are functions that operate on other functions, either by taking them as arguments or by returning them.

```go
package main

import "fmt"

func main() {

	igen := intGenerator()
	fmt.Println(">>", igen()) //2
	fmt.Println(">>", igen()) //4
	fmt.Println(">>", igen()) //6
	//igen = nil

}

type genType func() int

func intGenerator() genType {
	x := 0
	return func() int {
		x += 2
		return x
	}
}
```
## Lazy Evaluation
Lazy evaluation or non-strict evaluation is the process of delaying evaluation of an expression until it is needed. In general, Go does strict/eager evaluation but for operands like && and || it does a lazy evaluation. We can utilize higher-order-functions, closures, goroutines, and channels to emulate lazy evaluations.

## Lazy loading using sync.Once

```go

package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func Make(f func() int) LazyInt {
	var v int
	var once sync.Once
	return func() int {
		once.Do(func() {
			v = f()
			f = nil // so that f can now be GC'ed
		})
		return v
	}
}

func main() {
	n := Make(func() int { return 23 }) // Or something more expensive…
	fmt.Println(n())                    // Calculates the 23
	fmt.Println(n() + 42)               // Reuses the calculated value
}
```

## [Lazy loading implementation using higher order functions](https://go.dev/play/p/Ty8vpd52TaI)
```go
package main

import "fmt"

//   a func(x int) int
func Double(x int) int {
	return x * x
}

//    a func(x, y int) int
func Double2(x, y int) int {
	return x * y
}

//      Double2 == func(x int) int, x int) int
func ClosureFunc(a func(x, y int) int, x, y int) int {
	return a(x, y)
}

func main() {
	//fmt.Println(">>", ClosureFunc(Double, 5))
	fmt.Println(">>", ClosureFunc(Double2, 5, 6))  
}
```
## Futures in Go
Futures are mechanisms for decoupling a value from how it was computed. Goroutines and channels allow implementing futures trivially. 
```go
c := make(chan int)      // future
go func() { c <- f() }() // async
value := <-c             // await
```

## Singleton Pattern in Go or sync.Once
Just Call Your Code Only Once

```go
package main

import (
	"fmt"
	"sync"
)

type DbConnection struct{}

var (
	dbConnOnce sync.Once
	conn       *DbConnection
)

func GetConnection() *DbConnection {
	dbConnOnce.Do(func() {
		conn = &DbConnection{}
		fmt.Println("Inside")
	})
	fmt.Println("Outside")
	return conn
}

func main() {
	for i := 0; i < 5; i++ {
		_ = GetConnection()
	}
}
//outputs
Inside
Outside
Outside
Outside
Outside
Outside
```


## Resource
* [Lazy evaluation](https://deepu.tech/functional-programming-in-go/)
* [Go-Pipelines Lazy vs Eager](https://medium.com/@j.d.livni/understanding-go-pipelines-in-5-minutes-2906a5c41496)
* [lazy-evaluation-in-go](https://blog.merovius.de/posts/2015-07-17-lazy-evaluation-in-go/)
* [Futures in GO](https://appliedgo.net/futures/#:~:text=Futures%20in%20a%20nutshell,for%20Go's%20built%2Din%20concurrency.)
* [Functional-Programming](https://deepu.tech/functional-programming-in-go/)
* [Easy-functional-programming-techniques](https://dev.to/deepu105/7-easy-functional-programming-techniques-in-go-3idp)
* [Thread-safe-lazy-loading](https://www.fareez.info/blog/thread-safe-lazy-loading-using-sync-once-in-go/)
