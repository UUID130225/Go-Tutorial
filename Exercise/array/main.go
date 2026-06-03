package main

import (
	"fmt"
	"time"
)

var id int
var name string
var birt time.Time

type Student struct {
	StudentID int
	Fullname  string
	BirtDay   time.Duration
}



func add() bool {
	return true
}

func remove() bool {
	return true
}

func main() {
	// vie: Bài tập làm quyen với mảng
	fmt.Printf("Exercises to familiarize yourself with Array\n")

	for {
		var choice int;
		// vie: Mời bạn chọn menu
		fmt.Print("Please choose the menu: ")
		
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Please enter a number")
			continue
		}

		switch choice {
			case 1:
				fmt.Printf("Choice 1\n");
				
			case 2:
				fmt.Printf("Choice 2\n")
			case 0:
				return
			default:
				fmt.Println("Invalid choice ???")
		}
	}
	
}