package main

import (
	"fmt"
	"sync"
)

type post struct { // A post has a number of views
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) { // Increment the number of views on a post
	defer func() { // Use defer to ensure the mutex is unlocked even if a panic occurs
		p.mu.Unlock() // Unlock the mutex when done
		wg.Done()     // Signal that the goroutine is done
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup to wait for all goroutines to finish
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1) // Add a goroutine to the WaitGroup
		go myPost.inc(&wg)

	}

	wg.Wait() // Wait for all goroutines to finish
	// After all goroutines have finished, print the number of views
	fmt.Println(myPost.views)
}
