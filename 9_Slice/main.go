package main

import "fmt"

func main() {
	// slice -> dynamic array

	// var nums []int
	// fmt.Println(nums==nil)
	// fmt.Println(len(nums))

	nums := []int{}
	nums = append(nums,1)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))


}
