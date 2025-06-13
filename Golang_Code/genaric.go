package main

import (
	"fmt"
)

func generaric[T comparable](items []T) { // using generic we can use multiple type of slice int,string,bool etc
	for _, item := range items { // T any allows every type of slice so we use T int | string
		fmt.Println(item) //here it allows all three type of string
	}
}

// func stringslice(items []string) {     //in both the function logic is
// almost same so we use genarics
//
//		for _, item := range items {
//			fmt.Println(item)
//		}
//	}

type stack[T any] struct {
	elements []T
}

func main() {

	myStack := stack[string]{
		elements: []string{"golang"},
	}

	fmt.Println(myStack)
	// num := []int{1, 4, 5}

	// num := []string{"Golang", "Java", "Rust"}
	// // generaric(num)

}
