package main

import "fmt"

// add: phep cong
func add(x, y float64) float64 {
	return x + y
}

// subtract: Phep tru
func subtract(x, y float64) float64 {
	return x - y
}

// multiply: Phep nhan
func multiply(x, y float64) float64 {
	return x * y
}

// divide: Phep chia
func divide(x, y float64) float64 {
	return x / y
}

func main() {
	a := 30.0435
	b := 32.0

	fmt.Printf("%.4f + %.4f = %.4f\n", a, b, add(a, b))
	fmt.Printf("%.4f - %.4f = %.4f\n", a, b, subtract(a, b))
	fmt.Printf("%.4f * %.4f = %.4f\n", a, b, multiply(a, b))
	fmt.Printf("%.4f / %.4f = %.4f\n", a, b, divide(a, b))
}