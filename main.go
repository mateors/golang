package main

import "fmt"

type Nothing interface{}

func main() {

	var n Nothing

	fmt.Printf("%v %T\n", n, n)

	n = 1
	fmt.Printf("%v %T\n", n, n)

	n = "Hello"
	fmt.Printf("%v %T\n", n, n)

	n = 1.50
	fmt.Printf("%v %T\n", n, n)

}
