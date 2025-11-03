package patterns

import (
	"fmt"
	"sync"
)

const (
	totalJobs = 4
	workers   = 2
)

func WorkerPoolPattern() {
	jobs := make(chan int, totalJobs)
	result := make(chan int, workers)

	for i := 1; i <= workers; i++ {
		go Worker(i, jobs, result)
	}
	// Send jobs
	for j := 1; j <= totalJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// Receive results
	for a := 1; a <= totalJobs; a++ {
		<-result
	}

	close(result)
}

func Worker(id int, jobs <-chan int, results chan<- int) {

	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(job int) {
			defer wg.Done()
			fmt.Printf("woker : %d , is started the job : %d \n", id, job)
			result := job * 2
			results <- result
			fmt.Printf("woker : %d , is finished the job : %d \n", id, job)
		}(j)
	}

	wg.Wait()

}
