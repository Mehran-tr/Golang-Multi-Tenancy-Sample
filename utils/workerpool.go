package utils

import (
	"fmt"
	"time"
)

type Job struct {
	UserID    int
	UserEmail string
}

var JobQueue = make(chan Job, 100) // Buffer of 100 jobs

// Worker pool to process jobs concurrently
func Worker(id int, jobs <-chan Job) {
	for job := range jobs {
		fmt.Printf("Worker %d: Processing job for user %d\n", id, job.UserID)
		// Simulate sending an email
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d: Email sent to %s\n", id, job.UserEmail)
	}
}

func StartWorkerPool(workerCount int) {
	for i := 1; i <= workerCount; i++ {
		go Worker(i, JobQueue)
	}
}
