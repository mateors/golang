# Advance Golang

## FreeOSMemory
Go is a garbage collected language, which means that memory allocated and used by variables is automatically freed by the garbage collector when those variables become unreachable (if you have another pointer to the variable, that still counts as "reachable").

Freed memory does not mean it is returned to the OS. Freed memory means the memory can be reclaimed, reused for another variable if there is a need. So from the operating system you won't see memory decreasing right away just because some variable became unreachable and the garbage collector detected this and freed memory used by it.

The Go runtime will however return memory to the OS if it is not used for some time (which is usually around 5 minutes). If the memory usage increases during this period (and optionally shrinks again), the memory will most likely not be returned to the OS.

If you wait some time and not allocate memory again, freed memory will be returned to the OS eventually (obviously not all, but unused "big chunks" will be). If you can't wait for this to happen, you may call debug.FreeOSMemory() to force this behavior:

### How long should i wait to check how much memory is freed?
you should wait at least 7 minutes to check how much memory is freed. Sometimes it needs two GC passes, so it would be 9 minutes.
If that is not working, or it is too much time, you can add a periodic call to FreeOSMemory (no need to call runtime.GC() before, it is done by debug.FreeOSMemory() )

```go
package main

import (
    "runtime/debug"
    "time"
)

func main() {

    go periodicFree(1 * time.Minute)
    // Your program goes here
}

func periodicFree(d time.Duration) {
    tick := time.Tick(d)
    for _ = range tick {
        debug.FreeOSMemory()
    }
}

```

* [FreeOSMemory](https://pkg.go.dev/runtime/debug#FreeOSMemory) forces a garbage collection followed by an attempt to return as much memory to the operating system as possible.
* Keep in mind that it won't do anything unless the GC has been run
* The amount of memory released shown by runtime.ReadMemStats
 
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

## Reference
* https://stackoverflow.com/questions/18017979/golang-pointer-to-function-from-string-functions-name
* https://mikespook.com/2012/07/function-call-by-name-in-golang/

## Resource
* [Bitshifting Basic video](https://www.youtube.com/watch?v=qq64FrA2UXQ)
* https://github.com/yangchenxing/go-map2struct
* https://github.com/yangchenxing/go-singleflight
* [Map to Struct](https://github.com/mateors/go-map2struct)
* [free-memory-once-occupied](https://stackoverflow.com/questions/37382600/cannot-free-memory-once-occupied-by-bytes-buffer)
