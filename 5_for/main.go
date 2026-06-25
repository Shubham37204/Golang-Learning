package main

import "fmt"

func main() {

	//while loop  interm of for loop
	// i:=1
	// for i<4{
	// 	fmt.Println(i)
	// 	i=i+1
	// }

	// for i :=0; i<3; i++{
	// 	if i==2{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//using range in for loop
	for i := range 3 {
		fmt.Println(i)
	}

}
