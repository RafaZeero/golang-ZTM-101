package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

// Job: Receive channel / Read-only
// Results: Write channel / Write-only
// Control: Bidirectional / Read & Write
func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			default:
				panic("unhandled control message")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {
	// job
	job := make(chan Job, 50)
	// results
	results := make(chan Result, 50)
	// control
	control := make(chan ControlMsg)

	go doubler(job, results, control)

	for i := 0; i < 30; i++ {
		job <- Job{i}
	}

	for {
		select {
		case results := <-results:
			fmt.Println("results:", results)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			control <- DoExit // Send DoExit to control chan
			<-control         // Reading out the control chan
			fmt.Println("program exit")
			return
		}
	}
}
