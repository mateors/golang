# Advance Golang

## Topics
* Interface
* Pointer
* Singleflight | Mutex | Sync
* Bitwise operator | Left shift | right shift
* Goroutine
* Channel | Select
* Race condition | [Race tracer](https://pkg.go.dev/runtime/trace#Start)
* Context

> Package singleflight provides a duplicate function call suppression

> go doc & go tool usage details

## How do we call function by its name (string)

```golang

package main

import "fmt"

func someFunction1(a, b int) int {
        return a + b
}

func someFunction2(a, b int) int {
        return a - b
}

func someOtherFunction(a, b int, f func(int, int) int) int {
        return f(a, b)
}

func main() {
      
      //Method -1
        fmt.Println(someOtherFunction(111, 12, someFunction1))
        fmt.Println(someOtherFunction(111, 12, someFunction2))
        
      //Method-2  
      m := map[string]func(int, int) int {
              "someFunction1": someFunction1,
              "someFunction2": someFunction2,
      }
      z := someOtherFunction(x, y, m[key])
      fmt.Println(z)

}
```

* https://stackoverflow.com/questions/18017979/golang-pointer-to-function-from-string-functions-name

## Resource
* [Bitshifting Basic video](https://www.youtube.com/watch?v=qq64FrA2UXQ)
* https://github.com/yangchenxing/go-map2struct
* https://github.com/yangchenxing/go-singleflight
* [Map to Struct](https://github.com/mateors/go-map2struct)
