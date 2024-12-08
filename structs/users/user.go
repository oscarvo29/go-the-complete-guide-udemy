package users

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthData string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email string, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthData: "----",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("did not recieve a valid input")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthData: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthData)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
