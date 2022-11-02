

```go
package main

import "fmt"

func main() {

	//most repeated character
	//fmt.Println(90 - 65)  //A-Z
	//fmt.Println(122 - 97) //a-z

	input := "aAbcacdddaa"

	var char rune
	var freq int

	amap := make(map[rune]int)

	for _, c := range input {

		//fmt.Println(i, c)
		if c >= 97 && c <= 122 {
			fmt.Println("small:", c, string(c))

			if res, isOk := amap[c]; isOk {
				amap[c] = res + 1
			} else {
				amap[c] = 1
			}
		}
		if c >= 65 && c <= 90 {
			fmt.Println("CAPITAL:", c, string(c))
			if res, isOk := amap[c]; isOk {
				amap[c] = res + 1
			} else {
				amap[c] = 1
			}
		}

		count := amap[c]
		if count > freq {
			freq = count
			char = c
		}

	}

	fmt.Println(amap, char)

}

```
## Race condition

A race condition in Go occurs when two or more goroutines have shared data and interact with it simultaneously. 

```go

package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	var result []float64

	wg.Add(1)
	go func() {
		fmt.Println("worker 1")
		result = append(result, 50.50)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("worker 2")
		result = append(result, 78.50)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(result)
}
```
