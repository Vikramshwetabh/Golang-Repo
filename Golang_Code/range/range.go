package main

import "fmt"

//iterating for slice or array

func main() {
	//nums := []int{2, 3, 5, 6}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	//2nd method usine range keyword
	// for i, nums := range nums {
	// 	fmt.Println(i, nums)
	// }

	//Use of range on map

	m := map[string]string{"Fname": "John", "Lname": "Wick"}

	for k, v := range m { //k=key, v=value . we can return two values by using for loop in go
		fmt.Println(k, v)
	}
}
