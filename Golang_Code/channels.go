package main

import (
	"fmt"
	"time"
)

// read the channel
func processNum(numchan chan int) { //in a go routine numchan is passed of  type chan and chan is type of int
	for num := range numchan { //data is received
		fmt.Println(num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2

	result <- numResult //send
}

func emailsender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }() //synchronize the goroutine
	for email := range emailChan {
		fmt.Println("Sending the email to", email)

		time.Sleep(time.Second)
	}
}
func main() {
	// This is a simple example of using channels in Go
	chan1 := make(chan int)    //create a channel of type int
	chan2 := make(chan string) //create a channel of type string

	go func() {
		chan1 <- 10 //send data to channel

	}()
	go func() {
		chan2 <- "Hello" //send data to channel
	}()
	// The select statement is used to wait on multiple channel operations
	// It allows a goroutine to wait on multiple communication operations
	for i := 0; i < 2; i++ {
		select {
		case val1 := <-chan1: //receive data from channel
			fmt.Println("Received from chan1:", val1)
		case val2 := <-chan2: //receive data from channel
			fmt.Println("Received from chan2:", val2)
		default:

			fmt.Println("No data received from either channel") //if no data is received from both channels
		}

		// emailChan := make(chan string, 100) //buffered channel
		// done := make(chan bool)             //unbuffered channel
		// go emailsender(emailChan, done)     //start the goroutine
		// for i := 0; i < 5; i++ {
		// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
		// }
		// close(emailChan) //close the channel to avoid deadlock
		// fmt.Println("done sending")

		// <-done // blocks the channel until the goroutine finishes

		// result := make(chan int)
		// go sum(result, 5, 6)
		// res := <-result
		// fmt.Println(res) //received

		// numchan := make(chan int)
		// go processNum(numchan)
		// for {
		// 	numchan <- rand.Intn(100) //here data is send in queue
		// }

		// messageChain := make(chan string) //channel is created using chan keyword
		// messageChain <- "Hello" //data is send
		// msg := <-messageChain //data is recieved . Block channel
		// fmt.Println(msg) // this is situation of deadlock
	}
}
