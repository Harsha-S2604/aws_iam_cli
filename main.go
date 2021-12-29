package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Welcome to the AWS IAM Command line interface :)")
	fmt.Println("Please select a operation to perform")
	fmt.Println("========================================================")
	fmt.Println("1. Create a group\n2. Delete a group\n3. Create a user\n4. Delete a user")
	
	// get the input from user using Scanf function
	var option int
	fmt.Printf("Enter the number: ")
	fmt.Scanf("%d", &option)
	
	switch(option) {
		case 1:
		    CreateGroup()
		case 2:
		    DeleteGroup()
		case 3:
		    CreateUser()
		case 4:
		    DeleteUser()
	}
}

func CreateGroup() {
	fmt.Println("Create group")
}

func DeleteGroup() {
	fmt.Println("Delete group")
}

func CreateUser() {
	fmt.Println("Create user")
}

func DeleteUser() {
	fmt.Println("Delete user")
}
