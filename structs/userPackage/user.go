package userPackage

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	// age int   custom datatype new
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     // OR use "Superuser User"
	//Here "Superuser User" is doing Explicit embedding so we will use "User" to perform anonymous embedding
}

// Example of Struct embedding
func (anyUser *User) OutputUserDetails() { // Receiver argument function of struct you can user *user or user botha are just saves memory
	fmt.Println(anyUser.firstName, anyUser.lastName, anyUser.birthdate)
}

func (anyUser *User) ClearUsername() { //Pointer to edit actual value
	anyUser.firstName = ""
	anyUser.lastName = ""
}

func New(userFirstName, userLastName, userBirthdate string) (*User, error) { // utility function to assign values and reurn struct

	if userFirstName == "" || userLastName == "" || userBirthdate == "" {

		return nil, errors.New(" Firstname, Lastname and Birthdate is needed")
	}
	return &User{ // This is also called Struct literal or Composite literal
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email string, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New(" Email and Password is needed for Admin")
	}
	return &Admin{ // This is also called Struct literal or Composite literal
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "----",
			createdAt: time.Now(),
		},
	}, nil

	// return &Admin{ // This is also called Struct literal or Composite literal
	// 	email:    email,
	// 	password: password,
	// }, nil

}
