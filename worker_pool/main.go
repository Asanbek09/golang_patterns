package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("Worker: ", id, "Started job", j, "...")
		time.Sleep(time.Second)
		fmt.Println("Worker ", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numberOfjobs = 5

	jobs := make(chan int, numberOfjobs)
	results := make(chan int, numberOfjobs)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= numberOfjobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numberOfjobs; a++ {
		<-results
	}
}