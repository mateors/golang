
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	//wg.Add(1)

	c1 := make(chan map[string]string, 1) //no wait required
	go sendMap(c1)
	//go receiveMap(c1)               //finish this
	fmt.Println("Welcome", len(c1)) //then execute this line

	fmt.Println(<-c1)
	// for v := range c1 {
	// 	fmt.Println(v)
	// }
	//wg.Wait()

}

//send only or incoming
func sendMap(cstr chan<- map[string]string) {

	//defer wg.Done()
	time.Sleep(time.Millisecond * 4000)
	nmap := map[string]string{
		"name":      "Mostain",
		"education": "CIS",
	}
	cstr <- nmap
	close(cstr)

}

//receive only
func receiveMap(cstr <-chan map[string]string) {

	defer wg.Done()
	nmap := <-cstr
	for key, val := range nmap {
		fmt.Println(key, val)
	}

}

//send only
func send(cstr chan<- string) {

	time.Sleep(time.Millisecond * 100)
	cstr <- "Hello"

}

//receive only
func receive(cstr <-chan string) {

	msg := <-cstr
	fmt.Println(msg)

}
```

