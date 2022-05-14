# Break number into integer and fraction parts in Golang
math package of golang provides a Modf method that can be used to break a floating-point number into integer and floating part. Please note that the integer part is also returned as a float by this function.

```go
func Modf(f float64) (int float64, frac float64)
```

## [Test code](https://go.dev/play/p/z9--C7BGw1u)

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	floats := []float64{1.9999, 2.0001, 2.0}

	for _, f := range floats {

		//Contain both integer and fraction
		integer, fraction := math.Modf(f)
		fmt.Printf("Integer: %f. Fraction: %f\n", integer, fraction)

	}
}
```

## Output
```
Integer: 1.000000. Fraction: 0.999900
Integer: 2.000000. Fraction: 0.000100
Integer: 2.000000. Fraction: 0.000000
```
