package main

import "fmt"

const age = 30

func main() {
	//different ways to define constant
	
	//const age string = "golang"
	const name = "golang"

	//defining a constant block declaration
	const (
		port =5000
		host = "localhost"
	)

	fmt.Println(name,age)
	fmt.Println(port)
}
