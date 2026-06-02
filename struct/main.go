package main

import (
	"fmt"
)

type User struct {
	username string;
	password string;
}

func main() {
	var user1 User;

	user1.username = "nthung1325"
	user1.password = "admin123"

	fmt.Printf("User: %s - Password: %s ", user1.username, user1.password)
}