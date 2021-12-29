package main

import (
	"fmt"
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {

	// welcome note and description
	fmt.Println("Hello, Welcome to the AWS IAM Command line interface :)")

	// reads the access key from ~/.aws/config and ~/.aws/credentials
	cfg, cfgErr := config.LoadDefaultConfig(context.TODO())
	if cfgErr != nil {
		log.Fatal("configuration error. Please configure the AWS. For more information visit https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html")
	}

	iamClient := iam.NewFromConfig(cfg)
	
	fmt.Println("Please select any option to perform operation")
	fmt.Println("========================================================")
	fmt.Println("1. Create a group\n2. Delete a group\n3. Create a user\n4. Delete a user")
	
	// get the input from user using Scanf function
	var option int
	fmt.Printf("Enter the option: ")
	fmt.Scanf("%d", &option)
	
	// call function based on the user input
	/*
	   if user selects
	   1 => call CreateGroup function
	   2 => call DeleteGroup function
	   3 => call CreateUser function
	   4 => call  DeleteUser function
	*/
	switch(option) {
		case 1:
		    fmt.Println("\n=============== CREATE GROUP ===============")
		    CreateGroup(iamClient)
		case 2:
		    DeleteGroup(iamClient)
		case 3:
		    CreateUser(iamClient)
		case 4:
		    DeleteUser(iamClient)
	}
}

func CreateGroup(iamClient *iam.Client) {
	var groupName string
	fmt.Println("Enter a group name")
	fmt.Scanf("%s", &groupName)
	createGrpInp := &iam.CreateGroupInput{
		GroupName: aws.String(groupName),
	}
	
	createGroupResult, createGrpErr := iamClient.CreateGroup(context.TODO(),createGrpInp)
	if createGrpErr != nil {
		fmt.Println("Error in creating the group", createGrpErr.Error())
	} else {
		fmt.Printf("Group %s created successfully\n", *(createGroupResult.Group.GroupName))
	}
}

func DeleteGroup(iamClient *iam.Client) {
	fmt.Println("Delete group")
}

func CreateUser(iamClient *iam.Client) {
	fmt.Println("Create user")
}

func DeleteUser(iamClient *iam.Client) {
	fmt.Println("Delete user")
}
