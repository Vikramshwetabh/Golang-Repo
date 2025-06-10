package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 5
	// //Simple Switch
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 5:
	// 	fmt.Println("Five")
	// }

	//Multiple Switch Case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its Weekend")
	default:
		fmt.Println("Its Workday")
	}

}
