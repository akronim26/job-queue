package main

import (
	"fmt"
	"time"
)

// The definition of a job
type Job struct {
	Id int
	msg string
}

// The function to process a job
func (j Job) Process() {
	fmt.Println("Processing job #%d: %s", j.Id, j.msg)
	time.Sleep(1 * time.Second) // simulating the workload
	fmt.Println("Completed processing #%d", j.Id)
}