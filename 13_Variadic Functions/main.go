package main

import "fmt"

func sum(nums ...int) int {

	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	nums := []int{2, 5, 7, 3, 4}
	result := sum(nums...)
	fmt.Println(result)

}
