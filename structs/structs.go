package main

import (
	"fmt"

	"ovo/structs-udemy/users"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user, err := users.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
