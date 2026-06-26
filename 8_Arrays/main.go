package main

import "fmt"

func main() {
	
	// Integer array
	nums := [4]int{2, 3, 4, 5}

	// Boolean array
	flags := [4]bool{true, false, true, false}

	// String array
	names := [4]string{"Go", "Java", "Python", "C++"}


	fmt.Println("nums:", nums)
	fmt.Println("nums length:", len(nums))

	fmt.Println("flags:", flags)
	fmt.Println("flags length:", len(flags))

	fmt.Println("names:", names)
	fmt.Println("names length:", len(names))
}
