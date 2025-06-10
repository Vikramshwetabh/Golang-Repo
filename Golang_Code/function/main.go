package main

import "fmt"

func anotherfun() {
	fmt.Println("This is an other function")
}
func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("This is main function")
	anotherfun()
	fmt.Println("Sum of two number is", add(3, 7))
}
