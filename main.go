package main

import (
	"fmt"
	userMethods "go-learn/users"
)

func main() {

	users := userMethods.GetAllUsers()

	fmt.Println("Users:", users)

	getUser := userMethods.GetUser("Roberto")

	if getUser != nil {
		fmt.Println("SpecificUser:", getUser)
	} else {
		fmt.Println("User not found")
	}

	userMethods.UpdateUserAge("Roberto", 45)

	users = userMethods.GetAllUsers()

	fmt.Println("updatedUsers:", users)

	userMethods.DeleteUserByName("Roberto")

	users = userMethods.GetAllUsers()

	fmt.Println("updatedUsers:", users)

}

// Main will have user call services

// Create user service

// Create user struct and memory variable

// Call functions from main
