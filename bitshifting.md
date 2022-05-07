# Bit Shifting in Golang
The Go programming language supports several bitwise operators including the followings:

```
 &   bitwise AND
 |   bitwise OR
 ^   bitwise XOR
&^   AND NOT
<<   left shift
>>   right shift
```


## Left Shifting

```go
x := 7 //00000111
fmt.Printf("%08b\n", x)
fmt.Println(x, x<<1) //14
fmt.Println(x, x<<2) //28
fmt.Println(x, x<<3) //56
fmt.Println(x, x<<4) //112
  ```
  
  ## Manual Left shifting
  ```go
fmt.Printf("%d\n", 0b00000111) //7
fmt.Printf("%d\n", 0b00001110) //14
fmt.Printf("%d\n", 0b00011100) //28
fmt.Printf("%d\n", 0b00111000) //56
fmt.Printf("%d\n", 0b01110000) //112
fmt.Printf("%d\n", 0b11100000) //224
  ```
  
## Right Shifting

```go
fmt.Println(x, x>>1) //3  0b00000011
fmt.Println(x, x>>2) //1  0b00000001
fmt.Println(x, x>>3) //0  0b00000000
 ```
 
 ```go
 func add(a, b int64) int64 {

	for b != 0 {

		// common bits of a and b go to carry
		carry := a & b

		// xor - sum bits where one is not set
		a = a ^ b

		// shift carry by 1
		b = carry << 1
	}
	return a
}

func subtract(a, b int64) int64 {

	for b != 0 {

		// common bits of a and b go to borrow
		borrow := ^a & b

		// xor - sum bits where one is not set
		a = a ^ b

		// shift carry by 1
		b = borrow << 1
	}
	return a
}

 ```

## Resource
* [Golang IOTA & Bitwise Bit Shifting](https://www.youtube.com/watch?v=FHEyulwEEdE)
* [bit-manipulation-go](https://mariadesouza.com/2017/09/14/bit-manipulation-go)
