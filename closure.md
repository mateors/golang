# Golang Lazy Evaluation using closure
> What is a closure in computing? \
In computer science, a closure is a function that has an environment of its own. In this environment, there is at least one bound variable (a name that has a value, such as a number). The closure's environment keeps the bound variables in memory between uses of the closure.

## Higher-order function
In mathematics and computer science, a higher-order function is a function that does at least one of the following: takes one or more functions as arguments, returns a function as its result. All other functions are first-order functions. In mathematics higher-order functions are also termed operators or functionals.

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

## Resource
* [Lazy evaluation](https://deepu.tech/functional-programming-in-go/)
* [Go-Pipelines Lazy vs Eager](https://medium.com/@j.d.livni/understanding-go-pipelines-in-5-minutes-2906a5c41496)
* [lazy-evaluation-in-go](https://blog.merovius.de/posts/2015-07-17-lazy-evaluation-in-go/)
