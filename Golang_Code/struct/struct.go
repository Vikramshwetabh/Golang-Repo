package main

import (
	"fmt"
	"time"
)

type order struct { //intialising a struct
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func main() {
	myOrder := order{ //creating an instance of a struct just like creating an object
		id:     "1", //intialiing the value of data member of struch
		amount: 50.32,
		status: "delivered",
	}

	myOrder2 := order{
		id:        "2",
		amount:    86.33,
		status:    "Paid",
		createdAt: time.Now(),
	}

	myOrder.createdAt = time.Now()
	myOrder.status = "failed"  //->that 
	fmt.Println("Printing 2nd object", myOrder2)
	fmt.Println("Printing Structure", myOrder)
}
