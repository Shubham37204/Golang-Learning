package main

import "fmt"

func main() {

	//nums := []int{1,2,3}
	//sum :=0

	//using range in for loop
	// for i := range 3 {
	// 	fmt.Println(i)
	// }

	// i is the index or the key
	//num is value
	// for i,num:=range nums{
	//sum += num
	// 	fmt.Println(i,num)
	// }

	// marks := map[string]int{
	// 	"Math":    95,
	// 	"Science": 88,
	// 	"English": 90,
	// }
	// for k, v := range marks {
	// 	fmt.Println(k, v)
	//fmt.Println(v)
	// }

	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

	for i, c := range "golang" {
		fmt.Printf("Index = %d Character = %c ASCII = %d\n", i, c, c)
	}

}
