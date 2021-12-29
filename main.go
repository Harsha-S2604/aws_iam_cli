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
		    fmt.Println("create group")
		case 2:
		    fmt.Println("Delete group")
		case 3:
		    fmt.Println("Create user")
		case 4:
		    fmt.Println("Delete user")
	}
}
