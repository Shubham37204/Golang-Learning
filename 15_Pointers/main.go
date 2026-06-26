package main

import "fmt"

//Pass by Value
// func update(num int) {
// 	num = 100
// 	fmt.Println("Inside Function:", num)
// }

//pass br refrences
func update(num *int) {
	*num = 100
}

func main() {

	x := 10
	// update(x)

	update(&x)

	fmt.Println("Outside Function:", x)
}
