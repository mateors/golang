# Golang introduces Generics in version 1.18

## What is Generics?
Generics means parameterized types. The idea is to allow type (Integer, String, … etc., and user-defined types) to be a parameter to methods, classes, and interfaces. Using Generics, it is possible to create functions that work with different data types.

With generics, you can declare and use functions or types that are written to work with any of a set of types provided by calling code.

## Why should i care about generics?
You should care about generics because they mean you don’t have to write as much code, especially if you’re in the business of writing packages and tooling. It can be frustrating to write utility functions without generics support. Think about common data structures like binary search trees and linked lists. Why would you want to rewrite them for every type they could possibly contain? int, bool, float64, and string aren’t the end of the list, because you may want to store a custom struct type.

Generics will finally give Go developers an elegant way to write amazing utility packages.

## Generics 3 main features
* Type Parameter with cnstraint
* Type Inference
* Type Set

### 1. Type Parameter with constraints
```go
package main
import "fmt"

func myFunc[T any](a T) T {
	return a
}

func main() {
	a := myFunc[int](10)
	fmt.Println(a)
}

```
#### Calling above function
```
a:=myFunc[int](10)
fmt.Println(a)
```

### 2. Type Inference
Inference is using observation and background to reach a logical conclusion. You probably practice inference every day. For example, if you see someone eating a new food and he or she makes a face, then you infer he does not like it. Or if someone slams a door, you can infer that she is upset about something.

Inference: the process of inferring something. -> deduce or conclude -> remove type when calling a function

> Calling differently according to the inference theory.
```go

a:=myFunc(10)
fmt.Println(a)

//any value with int64 or underlying type int64 are allowed
//with the special symbol called tilde  ~
func nextNumber[T ~int64|float64](a T) T {
    return a+1
}

**Hints:** notice a tilde sign before int64

a:=nextNumber(5)
fmt.Println(a)

type superInt int64
var b superInt=7
a:=nextNumber(b) //if you remove tilde sign from the function defination it wouldn't work.
fmt.Println(a)
```

### 3. Type Set (Declare a type constraint| Custom type)

```go

type myType interface{
 ~float64 | int64 //Declare a union of int64 and float64 inside the interface.
}

func nextNumber[T myType](a T) T {
    return a+1
}

a:=nextNumber(7) 
fmt.Println(a)

type superInt float64
var b superInt=7
a:=nextNumber(7) 
fmt.Println(a)

```

You are only allowed to make your custom constraint using the following types. [why?](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#operations-based-on-type-sets)

* Ordered is a type constraint that matches any ordered type.
* An ordered type is one that supports the <, <=, >, and >= operators.

### Type lists
* ~int | ~int8 | ~int16 | ~int32 | ~int64 |
* ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
* ~float32 | ~float64 |
* ~string

### Generic Types vs Generic Function
So we know that we can write functions that use generic types, but what if we want to create a custom type that can contain generic types? For example, a slice of order-able objects. The new proposal makes this possible.

```go
type comparableSlice[T comparable] []T

func allEqual[T comparable](s comparableSlice[T]) bool {
    if len(s) == 0 {
        return true
    }
    last := s[0]
    for _, cur := range s[1:] {
        if cur != last {
            return false
        }
        last = cur
    }
    return true 
}

func main() {
    fmt.Println(allEqual([]int{4,6,2}))
    // false

    fmt.Println(allEqual([]int{1,1,1}))
    // true
}
```


```go

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

func main() {
    firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
    fmt.Println(firstInts, secondInts)
    // prints [0 1] [2 3]

    firstStrings, secondStrings := splitAnySlice([]string{"zero", "one", "two", "three"})
    fmt.Println(firstStrings, secondStrings)
    // prints [zero one] [two three]
}

```

## Builtin Constraints
* any
* comparable
* Parametric constraints or Custom constraints

### Comparable constraint
The comparable constraint is a predefined constraint as well, just like the any constraint. When using the comparable constraint instead of the any constraint, you can use the != and == operators within your function logic.

## Generics Golang Playground
* https://gotipplay.golang.org/


## Why Generics?
* https://go.dev/blog/why-generics
* https://go.dev/blog/survey2020-results
* https://qvault.io/golang/how-to-use-golangs-generics
 
## Resource
* https://go.dev/doc/tutorial/generics
