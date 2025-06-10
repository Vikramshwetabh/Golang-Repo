package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}
func main() {
	increament := counter() //it stores a function

	fmt.Println(increament()) //when we pass the variable as an argument
	//that sotres function
	fmt.Println(increament())
}
