package main

import "fmt"

func main() {
	sum1 := 1
	for ; sum1 < 3; {
		sum1 += sum1
	}

	fmt.Println(sum1)

}
