# Bit Shifting in Golang

## Left Shifting

	```go
  x := 7 //00000111
	fmt.Printf("%08b\n", x)
	fmt.Println(x, x<<1) //14
	fmt.Println(x, x<<2) //28
	fmt.Println(x, x<<3) //56
	fmt.Println(x, x<<4) //112
  ```
