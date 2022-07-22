package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func goroutine(wg *sync.WaitGroup, id int) {
	defer wg.Done() // Run wg.Done at end of function
	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Second * time.Duration(rand.Intn(5)+1)) // Sleep 1 - 5 seconds
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	rand.Seed(time.Now().UnixMilli()) // Random seed
	var wg sync.WaitGroup             // Create a waitgroup to manage the goroutines
	var numGoRoutines int = 5         // 5 goroutines
	wg.Add(numGoRoutines)             // Add 5 goroutines to waitgroup
	for i := 0; i < numGoRoutines; i++ {
		go goroutine(&wg, i+1) // Run each goroutine, the function runs wg.Done at the end
	}
	wg.Wait() // Wait for all threads to finish
	fmt.Println("Main thread finished")
}
