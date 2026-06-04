package main

import (
	"fmt"
	"os"
)

func main() {

	// Write to a file
	err := os.WriteFile("text.txt", []byte("Hello, World!"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully!")

	// Read the file
	contents, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:", string(contents))

	// Open the file for appending
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Lỗi mở file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nDòng mới")
	if err != nil {
		fmt.Println("Lỗi ghi thêm:", err)
		return
	}

	fmt.Println("Ghi thêm thành công")
}