package main

import "fmt"

//Struct (Blueprint)
type User struct {
	Name     string
	Age      int
	Relation string
}

//constructor
func newOrder(Name string,
	Age int,
	Relation string) *User {
	u := User{
		Name:     Name,
		Age:      Age,
		Relation: Relation,
	}
	return &u
}

//receiver type
//This is called a method because it has a receiver.
//Receiver (r *User) means
// r is the receiver variable (like this in Java or self in Python).
// *User means it receives a pointer to a User.
//use * when to modify the struct value
func (r *User) ChangeRelation(Relation string) {
	r.Relation = Relation
}

func main() {

	//1 way to create struct
	//Note: if we do not specify and field then the default value will be provided based on tht data type
	// u := User{
	// 	Name: "Shubham",
	// 	Age:  23,
	// }
	// u.Relation = "single"
	// fmt.Println(u)

	// u1 := User{
	// 	Name: "Harsh",
	// 	Age: 23,
	// }
	// u.Relation="Taken"
	// fmt.Println(u)
	// fmt.Println(u1)



	//use of receiver method
	// u := User{
	// 	Name:     "Shubham",
	// 	Age:      23,
	// 	Relation: "single",
	// }
	// u.ChangeRelation("taken")
	// fmt.Println(u)



	//accessing the constructor
	// u2 := newOrder("Shubham",
	// 	23,
	// 	"Single")
	// fmt.Println(u2)



	//using the struct only once
	u3 := struct {
		name    string
		isTaken bool
	}{"shubham", false}
	fmt.Println(u3)

}
