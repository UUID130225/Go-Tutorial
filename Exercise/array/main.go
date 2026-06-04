package main

import (
	"fmt"
)

type Student struct {
	StudentID int
	Fullname  string
	BirtDay   string
}

var students []Student

func add(student Student) error {
	students = append(students, student)
	return nil
}

func remove(id int) error {
	for i, s := range students {
		if s.StudentID == id {
			students = append(students[:i], students[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("student with ID %d not found", id)
}

func get() ([]Student, error) {
	return students, nil
}

func main() {
	fmt.Printf("Exercises to familiarize yourself with Array\n")

	for {
		var choice int
		fmt.Print("Please choose the menu: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Please enter a number")
			continue
		}

		switch choice {
		case 1:
			addStudent := Student{
				StudentID: len(students) + 1,
				Fullname:  "Nguyen Van A",
				BirtDay:   "2000-01-01",
			}
			err := add(addStudent)
			if err != nil {
				fmt.Println("Error adding student:", err)
			} else {
				fmt.Println("Added successfully")
				fmt.Println(students)
			}
		case 2:
			students, err := get()
			if err != nil {
				fmt.Println("Error getting students:", err)
			} else {
				fmt.Println("Get successfully")
				fmt.Println(students)
			}
		case 3:
			var id int
			fmt.Print("Enter student ID to remove: ")
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println("Please enter a valid number")
				continue
			}
			err = remove(id)
			if err != nil {
				fmt.Println("Error removing student:", err)
			} else {
				fmt.Println("Removed successfully")
			}
		case 0:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
