package main

import (
	"context"
	"fmt"
	"sync"
)

// Definition of a worker
type Worker struct {
	Id       int
	Context  context.Context
	JobQueue <-chan Job // read-only channel for jobs because workers just take up jobs (don't produce them)
	wg       *sync.WaitGroup
}

// Function to make worker process jobs
func (w *Worker) Start() {
	go func() {
		for {
			select {
			case <-w.Context.Done():
				return
			case job, ok := <-w.JobQueue:
				if !ok {
					fmt.Printf("Worked #%d: Job Queue closed\n", w.Id)
				}
				job.Process()
				w.wg.Done()
			}
		}
	}() // The `()` here calls this function immediately
}
