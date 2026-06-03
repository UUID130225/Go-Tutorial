package main

import "fmt"

func main() {
	// vie: Switch case cơ bản
	fmt.Printf("Basic switch case\n")

	var choice int
	// vie: Mời bạn chọn menu
	fmt.Print("Please choose the menu: ")

	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	switch choice {
		case 1:
			fmt.Printf("Choice 1\n")

		case 2:
			fmt.Printf("Choice 2\n")
		case 0:
			return
		default:
			fmt.Println("Invalid choice ???")
		}
}
