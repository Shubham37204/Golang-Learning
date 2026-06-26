package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguage() (string, string, bool) {
// 	return "golang", "python", true
// }

// func processIt(fn func(a int)int){
// 	fn(1)
// }

func processIt()func(a int)int{
	return func(a int)int{
		return 4
	}
}

func main() {
	//result := add(10, 15)

	//lang1, lang2, _ := getLanguage()

	// fn := func(a int)int{
	// 	return 2
	// }
	// processIt(fn)

	fn := processIt()
	fn(6)
	fmt.Println()

}
