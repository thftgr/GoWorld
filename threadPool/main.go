package main

import (
	"fmt"
	"time"
)

type WorkList []func()

func StartJobs(f []func(), workerCount int) {
	poolSize := len(f) // 대기열 크기

	jobs := make(chan func(), poolSize)
	results := make(chan bool, poolSize)

	for w := 0; w < workerCount; w++ {
		go func(jobs <-chan func(), results chan<- bool) {
			for j := range jobs {
				j()
				results <- true
			}
		}(jobs, results)
	}

	for j := 0; j < poolSize; j++ {
		jobs <- f[j]
	}
	close(jobs)

	for a := 0; a < poolSize; a++ {
		<-results
		//fmt.Println(len(jobs))
	}
}

func main() {
	w := WorkList{}

	for i := 0; i < 10; i++ {
		w = append(w, func() {
			time.Sleep(time.Second)
			fmt.Println("hello", time.Now().Nanosecond(), i)
		})
	}
	StartJobs(w, 10)

}
