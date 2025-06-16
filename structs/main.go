package main

import "fmt"

type user struct {
	firstName string
	lastName string
	contactDetails
}

type contactDetails struct {
	email string
	zipCode int
}

func main() {
	newUser := user {
		firstName: "Rafa",
		lastName: "Cabrera Romero",
		contactDetails: contactDetails{
			email: "rafa@testestest.com",
			zipCode: 12345,
		},
	}
	newUser.updateUser()
	newUser.printUser()
}

func (u *user) updateUser() {
	(*u).firstName = "Pepe"
}

func (u user)printUser() {
	fmt.Printf("%-v", u)
}
