package main

import "fmt"

func main() {
	// //declaration of map
	// m := make(map[string]string) //here first string is type of key and
	// //  second is type of its value

	// //adding element in map
	// m["Name"] = "Shwetabh"
	// m["State"] = "Bihar"

	// fmt.Println(m)
	// delete(m, "State")//deletion of key in string
	// fmt.Println(m)

	//Another way of declearing map
	// m := map[string]int{"Mobile": 4, "Phone": 3}
	// fmt.Println(m)

	//To check either values is present or not
	m := map[string]int{"Mobile": 4, "Phone": 3}

	_, ok := m["Phone"] //here _(underscore) is a extra value thats return in golang

	if ok {
		fmt.Println("All good")
	} else {
		fmt.Println("Not ok")
	}

}
