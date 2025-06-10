package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In change num position", num)

// }

//by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}
func main() {
	num := 1
	// changeNum(num)
	changeNum(&num)

	// fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main", num)

}
