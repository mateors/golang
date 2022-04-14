# Golang Lazy Evaluation using closure
> What is a closure in computing? \
In computer science, a closure is a function that has an environment of its own. In this environment, there is at least one bound variable (a name that has a value, such as a number). The closure's environment keeps the bound variables in memory between uses of the closure.

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
