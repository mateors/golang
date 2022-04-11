
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

```go
package main

import (
	"fmt"
	"time"
)

func main(){

rows := []map[string]interface{}{
		{"name": "Mostain", "age": 38},
		{"name": "Moaz", "age": 25},
		{"name": "Nahid", "age": 26},
		{"name": "Anonnya", "age": 27},
		{"name": "Tareq", "age": 29},
		{"name": "Eshita", "age": 27},
		{"name": "Riyaz", "age": 21},
		{"name": "Pallabi", "age": 21},
		{"name": "Zubair", "age": 22},
		{"name": "Sumi", "age": 30},
	}
	c1 := make(chan []map[string]interface{}, 5)

	se := calcMaster(10, 2)
	for _, val := range se {
		go rowProcessor(c1, val, rows)

	}

	//time.Sleep(time.Millisecond * 100)
	var nmap = make([]map[string]interface{}, 0)
	for range se {
		vals := <-c1
		for _, row := range vals {
			nmap = append(nmap, row)
		}
		fmt.Println("receiver:", len(vals))
	}

	close(c1)
	fmt.Println("Done!", len(nmap))
}
	
//sender
func rowProcessor(c chan<- []map[string]interface{}, sval *startEnd, rows []map[string]interface{}) {

	fmt.Println("sender:", sval.Start, sval.End)
	//time.Sleep(time.Millisecond * 5000)
	c <- rows[sval.Start:sval.End]
}

type startEnd struct {
	Start int
	End   int
}
```

