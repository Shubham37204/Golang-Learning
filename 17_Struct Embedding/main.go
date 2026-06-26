package main

import "fmt"

type customer struct {
	phone int
	email string
}

type User struct {
	Name     string
	Age      int
	Relation string
	customer
}

func main() {

	// cust := customer{
	// 	phone: 4566123789,
	// 	email: "abc@gmail.com",
	// }

	u := User{
		Name:     "Shubham",
		Age:      23,
		Relation: "single",

		//customer: cust,

		customer: customer{
			phone: 4566123789,
			email: "abc@gmail.com",
		},
	}
	
	fmt.Println(u)

}
