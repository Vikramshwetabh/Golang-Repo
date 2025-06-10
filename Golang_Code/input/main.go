package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("What is your name")
	//fmt.Scan(&name)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("My name is", name)
}
