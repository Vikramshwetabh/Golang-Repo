package main

import "fmt"

func main() {
	//unintialize slice is nil
	// var nums []int
	// //fmt.Println(len(nums))
	// fmt.Println(nums == nil)

	// var nums = make([]int, 0, 5) //make method is used to intialize the array with zero

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// fmt.Println(nums)

	//copy one slice to another

	var nums = make([]int, 0)
	nums = append(nums, 2)
	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	fmt.Println(nums, nums2)

}
