package main

import (
	"fmt"
	"maps"
)

func main() {

	// make(map[keyType]valueType)

	//Creating a Map using Literal
	// marks := map[string]int{
	// 	"Math":    95,
	// 	"Science": 88,
	// 	"English": 90,
	// }
	// fmt.Println(marks)

	marks := make(map[string]int)

	//Adding Elements
	marks["Math"] = 95
	marks["Science"] = 88
	marks["English"] = 90

	fmt.Println(marks)

	//Accessing Values
	fmt.Println(marks["Math"])

	//Updating Values
	marks["Math"] = 100
	fmt.Println(marks)

	//Deleting Elements
	delete(marks, "Science")
	//clear(marks)
	fmt.Println(marks)

	//Length of a Map
	fmt.Println(len(marks))

	//1. Checking if a Key Exists
	value, exists := marks["Math"]
	fmt.Println(value)
	fmt.Println(exists)

	//2.other way to check if a Key Exists
	m := map[string]int{"price": 40, "phones": 3}
	v, ok := m["phones"]
	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	//Looping Through a Map
	for key, value := range marks {
		fmt.Println(key, value)
	}

	//checking if 2 map are wqual are not
	m1 := map[string]int{"price": 40, "phones": 3}
	fmt.Println(maps.Equal(m,m1))

}
