package main

import (
	"fmt"
	"context"
)

// Definition of a worker
type Worker struct {
	Id int
	Context context.Context
	JobQueue <-chan Job // read-only channel for jobs because workers just take up jobs (don't produce them)
}

// Function to make worker process jobs
func (w *Worker) Start() {
	go func() {
		fmt.Println("Worker #%d started processing!", w.Id)
		for {
			select {
			case <- w.Context.Done() :
				fmt.Println("Worker #%d stopping...", w.Id)
				return
			case job, ok := <- w.JobQueue:
				if !ok {
					fmt.Printf("Worked #%d: Job Queue closed", w.Id)
				}
				job.Process()
			}
		}
	}()
}

