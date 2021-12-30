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
	
	iamForLoop:for {
	fmt.Println("Please select any option to perform operation")
	fmt.Println("========================================================")
	fmt.Println("1. Create a group\n2. Delete a group\n3. Create a user\n4. Delete a user\n5. List groups\n6. List users\n7. Exit\n")
	
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
		case 5:
		    ListGroups(iamClient)
		case 6:
		    ListUsers(iamClient)
		case 7:
		    break iamForLoop
	}
	}
}

func CreateGroup(iamClient *iam.Client) {
	var groupName string
	fmt.Print("Enter a group name: ")
	fmt.Scanf("%s", &groupName)
	createGrpInp := &iam.CreateGroupInput{
		GroupName: aws.String(groupName),
	}
	
	createGroupResult, createGrpErr := iamClient.CreateGroup(context.TODO(),createGrpInp)
	if createGrpErr != nil {
		fmt.Println("Error in creating the group", createGrpErr.Error())
	} else {
		fmt.Printf("--------------- Group %s created successfully ---------------\n", *(createGroupResult.Group.GroupName))
	}
}

func DeleteGroup(iamClient *iam.Client) {
	fmt.Println("\n=============== DELETE GROUP ===============")
	ListGroups(iamClient)
	var groupName string
	fmt.Print("Enter the group name: ")
	fmt.Scanf("%s", &groupName)
	
	deleteGrpInp := &iam.DeleteGroupInput{
		GroupName: aws.String(groupName),
	}

	_, deleteGroupErr := iamClient.DeleteGroup(context.TODO(), deleteGrpInp)

	if deleteGroupErr != nil {
		fmt.Println("Error deleting the group", deleteGroupErr.Error())
	} else {
		fmt.Println("---------------- Successfully deleted the group ----------------\n")
	}
}

func CreateUser(iamClient *iam.Client) {
	fmt.Println("\n=============== CREATE USER ===============")
	var userName string
	fmt.Print("Enter the username to create: ")
	fmt.Scanf("%s", &userName)

	createUserInp := &iam.CreateUserInput {
		UserName: aws.String(userName),
	}

	createUserResult, createUserErr := iamClient.CreateUser(context.TODO(), createUserInp)

	if createUserErr != nil {
		fmt.Println("Error creating the user: ", createUserErr.Error())
	} else {
		fmt.Printf("---------------- User %s created successfully ----------------\n\n", *(createUserResult.User.UserName)) 
	}
}

func DeleteUser(iamClient *iam.Client) {
	fmt.Println("\n=============== DELETE USER ===============")
	ListUsers(iamClient)

	var userName string
	fmt.Print("Enter username to delete: ")
	fmt.Scanf("%s", &userName)

	deleteUserInp := &iam.DeleteUserInput {
		UserName: aws.String(userName),
	}

	_, deleteUserErr := iamClient.DeleteUser(context.TODO(), deleteUserInp)

	if deleteUserErr != nil {
		fmt.Printf("Error deleting the user: %s\n\n", deleteUserErr.Error())
	} else {
		fmt.Println("\n--------------- USER DELETED SUCCESSFULLY ----------------\n\n")
	}

}

func ListGroups(iamClient *iam.Client) {
	listGrpsInp := &iam.ListGroupsInput{}
	getGroupsResult, getGroupsErr := iamClient.ListGroups(context.TODO(), listGrpsInp)
	if getGroupsErr != nil {
		fmt.Println("Error fetching groups", getGroupsErr.Error())
	} else {
		groupsArr := getGroupsResult.Groups
		if len(groupsArr) > 0 {
			fmt.Println("\n=============== Here are the list of groups ================")
			for i := 0; i < len(groupsArr); i++ {
				fmt.Println(*(groupsArr[i].GroupName))
			}
		} else {
			fmt.Println("There are no groups in this account\n")
		}
	}
	fmt.Println()
}

func ListUsers(iamClient *iam.Client) {
	listUsersInp := &iam.ListUsersInput{}
	getUsersResult, getUsersErr := iamClient.ListUsers(context.TODO(), listUsersInp)
	if getUsersErr != nil {
		fmt.Println("Error fetching users", getUsersErr.Error())
	} else {
		usersArr := getUsersResult.Users
		if len(usersArr) > 0 {
			fmt.Println("\n=============== Here are the list of users ===============")
			for i := 0; i < len(usersArr); i++ {
				fmt.Println(*(usersArr[i].UserName) + " created on", *(usersArr[i].CreateDate))
			}
			fmt.Println()
		} else {
			fmt.Println("There are no users in this account\n")
		}
	}
}
