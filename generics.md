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

### Type Parameter with cnstrain
func myFunc[T any](a T) T {
    return a
} 

func myFunc[T interface{}](a T) T {
    return a
}

> Calling above function
a:=myFunc[T int](10)
fmt.Println(a)

### Type Inference
Inference is using observation and background to reach a logical conclusion. You probably practice inference every day. For example, if you see someone eating a new food and he or she makes a face, then you infer he does not like it. Or if someone slams a door, you can infer that she is upset about something.

Inference: the process of inferring something. -> deduce or conclude -> remove type when calling a function

> Calling differently according to the inference theory.
a:=myFunc(10)
fmt.Println(a)


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
## Builtin constraints
* any
* comparable
* Parametric constraints or Custom constraints

## Generics Golang Playground
* https://gotipplay.golang.org/


## Why Generics?
* https://go.dev/blog/why-generics
* https://go.dev/blog/survey2020-results
* https://qvault.io/golang/how-to-use-golangs-generics
 
## Resource
* https://go.dev/doc/tutorial/generics
