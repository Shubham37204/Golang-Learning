package main

// import (
// 	"fmt"
// 	"slices"
// )

func main() {

	// slice -> dynamic array

	//1.an empty slice
	// var nums []int
	// fmt.Println(nums==nil)
	// fmt.Println(len(nums))

	//2.Using a Slice Literal
	// arr := []int{}
	// arr = append(arr, 1)
	// fmt.Println(arr)

	//3.Using the make Function
	// nums := make([]int, 3, 5) or  nums := make([]int, 0, 5)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	//nums := make([]int, 3, 5)
	// nums[0] = 3
	// nums[1] = 10
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	//copy function
	// nums = append(nums, 2)
	// nums1 := make([]int, len(nums))
	// copy(nums1,nums)
	// fmt.Println(nums, nums1)

	//slice operator
	// var nums =[]int{1,2,3}
	// fmt.Println(nums[0:2])

	//slice
	// var nums1 = []int{1,2}
	// var nums2 = []int{1,2}
	// fmt.Println(slices.Equal(nums1,nums2))

	//2D slice
	//var nums=[][]{{1,2,3},{4,5,6}}
	//fmt.Println(nums)

}
