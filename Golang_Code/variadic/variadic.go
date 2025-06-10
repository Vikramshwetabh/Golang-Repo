package main

import "fmt"

func add(nums ...int) int {
	total := 0
	for _, num := range nums { //_(underscore)->index(i) and num is the value at that index
		total = total + num
	}

	return total
}
func main() {
	nums := []int{2, 3, 46, 774, 2, 4, 5}
	//result := add(2, 3, 46, 774, 2, 4, 5)
	result := add(nums...) //this is an other way of doing same thing
	fmt.Println(result)
}
