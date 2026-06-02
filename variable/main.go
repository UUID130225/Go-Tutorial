package main

import (
	"fmt"
)

var a string = "String a"
var b int = 1

func main() {
	c := "Go"
	
	fmt.Printf("%s\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%s\n", c)

	var (
		d string = "Go"
		e int = 5
	)
	fmt.Printf("%s\n", d)
	fmt.Printf("%d\n", e)
}