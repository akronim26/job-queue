package main

import (
	"context"
	"sync"
)

// Definition of a Dispatcher (one who manages pool of workers and job queue)
type Dispatcher struct {
	WorkerCount int
	JobQueue    chan Job
	workers     []*Worker
	ctx         context.Context
	cancel      context.CancelFunc
	wg          *sync.WaitGroup
}

// Function to create and initialise a new dispatcher
func NewDispatcher(workerCount, queueSize int) *Dispatcher {
	ctx, cancel := context.WithCancel(context.Background())

	return &Dispatcher{
		WorkerCount: workerCount,
		JobQueue:    make(chan Job, queueSize),
		ctx:         ctx,
		cancel:      cancel,
		wg:          &sync.WaitGroup{},
	}
}

// Function to start the job queue
func (d *Dispatcher) Start() {
	for i := 1; i <= d.WorkerCount; i++ {
		worker := &Worker{
			Id:       i,
			Context:  d.ctx,
			JobQueue: d.JobQueue,
			wg: d.wg,
		}
		d.workers = append(d.workers, worker)
		worker.Start()
	}
}

// Function to submit a job
func (d *Dispatcher) Submit(job Job) {
	d.wg.Add(1)
	go func() {
		d.JobQueue <- job
	}()
}

// Function to stop the dispatcher gracefully
func (d *Dispatcher) Stop() {
	d.cancel()
	d.wg.Wait()
	close(d.JobQueue)
}
