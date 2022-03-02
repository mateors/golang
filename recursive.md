## Recursive function
Best real world recurive function example

```go
func colid(num int) string {

	var alphabets string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// val := fmt.Sprintf(`%c`, alphabets[num-1])
	// fmt.Println(val)
	if num <= 0 {
		return ""
	}
	if num < 26 {
		return fmt.Sprintf(`%c`, alphabets[num-1])
	}
	rem := num % 26 //remainder
	fld := num / 26 //floor-division
	if rem == 0 {
		if fld == 1 {
			if rem == 0 {
				rem = 26
			}
			return fmt.Sprintf(`%c`, alphabets[rem-1])
		}
		if rem == 0 {
			rem = 26
		}
		return fmt.Sprintf(`%s%c`, colid(fld-1), alphabets[rem-1])
	}
	return fmt.Sprintf(`%s%c`, colid(fld), alphabets[rem-1])
}
```
usage:
fmt.Println(colid(1)) //A
fmt.Println(colid(26)) //Z
fmt.Println(colid(27)) //AA
fmt.Println(colid(52)) //AZ
fmt.Println(colid(705)) //AAC
