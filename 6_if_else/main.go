package main

import "fmt"

func main() {

	//age := 14
	// if age>=19{
	// 	fmt.Println("person is adult")
	// }else{
	// 	fmt.Println("person is not adult")
	// }

	// if age >= 19 {
	// 	fmt.Println("person is adult")
	// } else if age > 15 {
	// 	fmt.Println("person is teenager")
	// } else {
	// 	fmt.Println("person is not adult")
	// }

	//defining age inside of if statement
	if age := 15; age >= 18 {
		fmt.Println("person is adult")
	} else if age >= 12 {
		fmt.Println("person is adult teenager")
	} else {
		fmt.Println("person is not adult")
	}
}
