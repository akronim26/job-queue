package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	workerCount := 4
	queueSize := 10

	fmt.Println("Starting dispatcher...")
	dispatcher := NewDispatcher(workerCount, queueSize)
	dispatcher.Start()

	go func() {
		for i := 1; i <= 20; i++ {
			job := Job{
				Id:  i,
				Msg: fmt.Sprintf("Job Id: %v", i),
			}

			dispatcher.Submit(job)

		}

		time.Sleep(500 * time.Millisecond)
	}()

	// This creates a signal channel for receiving OS signals regarding termination (Ctrl+C, kill, Docker container stop)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan // Wait for the termination
	fmt.Println("\nStopping dispatcher...")
	dispatcher.Stop()
}
