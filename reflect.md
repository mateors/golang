# Reflect

## Array value
```go
	stds := []string{"Istiyak", "Hasan", "Dipu"}
	fmt.Println("kind:", reflect.ValueOf(stds).Kind()) //slice
	fmt.Println("type:", reflect.ValueOf(stds).Type())
	fmt.Println("length:", reflect.ValueOf(stds).Len())
	len := reflect.ValueOf(stds).Len()
	for i := 0; i < len; i++ {
		fmt.Println(i, "->", reflect.ValueOf(stds).Index(i))
	}
  ```
