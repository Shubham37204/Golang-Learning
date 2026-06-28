package main

import (
	"fmt"
	"my-awesome-project/auth"
	userdetails "my-awesome-project/userDetails"
)

func main() {

	auth.LoginCred("shubham", "pass123")
	sessionsVar := auth.SessionAuth()
	fmt.Print(sessionsVar)
	userdet := userdetails.User{
		Email: "shubha@gmail.com",
		Name: "Shubham",
	}
	fmt.Println(userdet)

}
