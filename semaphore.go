/// batch processing in go lang
var files []string // example list file path


sem := make(chan struct{}, 5) // put the number of batch here,, example 5
var wg sync.WaitGroup
for _, file := range files {
	wg.Add(1)
	//fmt.Println(file)
	sem <- struct{}{}
	go func() {
		defer wg.Done()
		defer func() { <-sem }()

		cmd := exec.Command("tesseract", file, "-")
		scan, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(scan))

	}()
}
wg.Wait()




///////////////////////////////////////////////////// OTHER EXAMPLE NOT TESTED/////////////////////

package main

import (
	"fmt"
	"sync"
	"time"
)

// Semaphore controls access to a limited number of resources.
type Semaphore struct {
	sem chan struct{}
}

// NewSemaphore creates a new semaphore with a given capacity.
func NewSemaphore(capacity int) *Semaphore {
	return &Semaphore{
		sem: make(chan struct{}, capacity),
	}
}

// Acquire acquires a semaphore permit, blocking if necessary.
func (s *Semaphore) Acquire() {
	s.sem <- struct{}{}
}

// Release releases a semaphore permit.
func (s *Semaphore) Release() {
	<-s.sem
}

func main() {
	// Number of items to process
	numItems := 20

	// Batch size (number of concurrent workers)
	batchSize := 5

	// Create a semaphore with the batch size as capacity
	sem := NewSemaphore(batchSize)

	var wg sync.WaitGroup

	// Simulate processing items in batches
	for i := 0; i < numItems; i++ {
		wg.Add(1)
		sem.Acquire() // Acquire a permit before starting a worker
		go func(item int) {
			defer wg.Done()
			defer sem.Release() // Release the permit when worker is done
			
			// Simulate work
			fmt.Printf("Processing item: %d\n", item)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println("All items processed.")
}
