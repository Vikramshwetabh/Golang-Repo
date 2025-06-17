package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) { // wait group pointer is passed
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1) //1 is added to wait the process
		go task(i, &wg)
	}

	wg.Wait()
}
