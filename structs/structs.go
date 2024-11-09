package main

import (
	"fmt"

	"example.com/structs/userPackage"
)

// type user struct {
// 	firstName string
// 	lastName  string
// 	birthdate string
// 	// age int   custom datatype new
// 	createdAt time.Time
// }

// func (anyUser user) outputUserDetails() { // Receiver argument fuccntion of struct
// 	fmt.Println(anyUser.firstName, anyUser.lastName, anyUser.birthdate)
// }

// func (anyUser *user) clearUsername() { //Pointer to edit actual value
// 	anyUser.firstName = ""
// 	anyUser.lastName = ""
// }

// func newUser(userFirstName, userLastName, userBirthdate string) (*user, error) { // utility function to assign values and reurn struct

// 	if userFirstName == "" || userLastName == "" || userBirthdate == "" {

// 		return nil, errors.New(" Firstname, Lastname and Birthdate is needed")
// 	}
// 	return &user{ // This is also called Struct literal or Composite literal
// 		firstName: userFirstName,
// 		lastName:  userLastName,
// 		birthdate: userBirthdate,
// 		createdAt: time.Now(),
// 	}, nil
// }

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *userPackage.User

	appUser, err := userPackage.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	//Another way

	// appUser = user {  // Here order should be same
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }
	// outputUserDetails(&appUser)

	admin, err := userPackage.NewAdmin("test@gam.com", "1234")

	if err != nil {
		fmt.Println(err)
		return
	}

	//  Needed this when explicit embedding is perfomed
	// admin.Superuser.OutputUserDetails()
	// admin.Superuser.ClearUsername()
	// admin.Superuser.OutputUserDetails()

	// Anonymous embedding
	admin.OutputUserDetails()
	admin.ClearUsername()
	admin.OutputUserDetails()
	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()

	// Practicing dynam,ic value assignmnet for admin variable
	// adminCred, err := userPackage.New(userFirstName, userLastName, userBirthdate)
	// admin.Superuser = *adminCred
	// admin.Superuser.OutputUserDetails()

	// (*appUser).outputUserDetails() // Now a method
	// // (*appUser).outputUserDetails()  a way to access as pointer
	// appUser.clearUsername()
	// appUser.outputUserDetails()

}

// func outputUserDetails(anyUser user) {
// 	fmt.Println(anyUser.firstName, anyUser.lastName, anyUser.birthdate)
// }

// func outputUserDetails(anyUser *user) {
// 	fmt.Println(anyUser.firstName, anyUser.lastName, anyUser.birthdate) // Go allows this
// 	// fmt.Println((*anyUser).firstName, (*anyUser).lastName, (*anyUser).birthdate) Technically correct way
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
