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
)

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
    defer wg.Done()

    // Acquire semaphore
    sem <- struct{}{}
    fmt.Printf("Worker %d started\n", id)

    // Do some work
    // ...

    // Release semaphore
    <-sem
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    maxConcurrency := 3
    sem := make(chan struct{}, maxConcurrency)
    var wg sync.WaitGroup

    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go worker(i, sem, &wg)
    }

    wg.Wait()
}


////////////////////////////////////////// BEST way inside the go func /////////////////


package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    maxGoroutines := 5
    semaphore := make(chan struct{}, maxGoroutines)

    var wg sync.WaitGroup
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            semaphore <- struct{}{}
            defer func() { <-semaphore }()
            
            // Simulate a task
            fmt.Printf("Running task %d\n", i)
            time.Sleep(2 * time.Second)
        }(i)
    }
    wg.Wait()
}
