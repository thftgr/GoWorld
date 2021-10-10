package src

import "fmt"

type WorkList []func()

func worker(jobs <-chan func(), results chan<- bool) {
	for j := range jobs {
		j()
		results <- true
	}
}

func StartJobs(f []func(), workerCount int) {
	poolSize := len(f) // 대기열 크기

	jobs := make(chan func(), poolSize)
	results := make(chan bool, poolSize)

	for w := 0; w < workerCount; w++ {
		go worker(jobs, results)
	}

	for j := 0; j < poolSize; j++ {
		jobs <- f[j]
	}
	close(jobs)

	for a := 0; a < poolSize; a++ {
		<-results
		fmt.Println(len(jobs))
	}
}
func (w *WorkList) AddJob(f func()) {
	*w = append(*w, f)
}
