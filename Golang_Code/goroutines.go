package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) { //this is inline function
			fmt.Println(i)
		}(i) //inline function call . here task runs parallelly not concurently
	}
	time.Sleep(time.Second * 2)
}
