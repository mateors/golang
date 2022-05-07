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
  
  ## Manual Left shifting
  ```go
fmt.Printf("%d\n", 0b00000111) //7
fmt.Printf("%d\n", 0b00001110) //14
fmt.Printf("%d\n", 0b00011100) //28
fmt.Printf("%d\n", 0b00111000) //56
fmt.Printf("%d\n", 0b01110000) //112
fmt.Printf("%d\n", 0b11100000) //224
  ```
  
