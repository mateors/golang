package main

import (
	_ "embed"
	"fmt"
)

//go:embed students.txt
var f string

func main() {

	//var f embed.FS
	//f.ReadDir()
	fmt.Println(f)

}
