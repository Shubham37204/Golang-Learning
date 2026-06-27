package main
import "fmt"

func printSlice[T any](items []T){
	for _,item := range items{
		fmt.Println(item)
	}
}


// func printSlice[T comparable](items []T){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T comparable, V string](items []T,name V){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T int | string | bool](items []T){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }


// type stack[T any] struct{
// 	elements []T
// }

func main() {

	//names :=[]string{"shubham"}

	// vals :=[]bool{true,false}
	// printSlice(vals)

	// myStack := stack[int]{
	// 	elements: []int{1,2,3,4},
	// }
	// fmt.Println(myStack)
}
