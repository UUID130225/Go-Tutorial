package main

import "fmt"

func main() {
	arr := [...]string{"hung", "Minh"}
	slide := []string{"hung", "Minh"}

	fmt.Printf("leng array %d", len(arr))
	fmt.Printf("leng slide %d", len(slide))

	fmt.Printf("%T", arr)
	fmt.Printf("%T", slide)
}