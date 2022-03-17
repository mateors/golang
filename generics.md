# Golang introduces Generics in version 1.18

## What is Generics?
Generics means parameterized types. The idea is to allow type (Integer, String, … etc., and user-defined types) to be a parameter to methods, classes, and interfaces. Using Generics, it is possible to create functions that work with different data types.

With generics, you can declare and use functions or types that are written to work with any of a set of types provided by calling code.

## Why should i care about generics?
You should care about generics because they mean you don’t have to write as much code, especially if you’re in the business of writing packages and tooling. It can be frustrating to write utility functions without generics support. Think about common data structures like binary search trees and linked lists. Why would you want to rewrite them for every type they could possibly contain? int, bool, float64, and string aren’t the end of the list, because you may want to store a custom struct type.

Generics will finally give Go developers an elegant way to write amazing utility packages.

```go

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
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
