package main

import (
	"fmt"
)

//ExecuteFunction as a goroutine that receives a task from a channel and writes output into another channel
func ExecuteFunction(id int, jobs <-chan func() error, results chan<- error, done chan<- bool) {
	for job := range jobs {
		results <- job()
	}
	done <- true
}

//Run a slice of functions using N goroutines and accept maximum of M errors from functions
func Run(tasks []func() error, N int, M int) error {
	numJobs := len(tasks)
	jobs := make(chan func() error, numJobs)
	results := make(chan error, numJobs)
	done := make(chan bool)
	errorCounter := 0
	completedTasks := 0

	for i := 0; i < N; i++ {
		go ExecuteFunction(i, jobs, results, done)
	}
	for len(tasks) != 0 && errorCounter < M {
		taskNum := N
		if len(tasks) < N {
			taskNum = len(tasks)
		}
		for i := 0; i < taskNum; i++ {

			jobs <- tasks[len(tasks)-1]
			tasks = tasks[:len(tasks)-1]
			completedTasks++
			err := <-results
			if err != nil {
				errorCounter++
			}
		}
	}
	close(jobs)
	for i := 0; i < N; i++ {
		<-done
	}
	if errorCounter >= M {
		return fmt.Errorf("too many errors, completed tasks: %d", completedTasks)
	}
	return nil
}

func ef1() error {
	return fmt.Errorf("error ef1")
}

func f1() error {
	return nil
}

func main() {
	fSlice := []func() error{ef1, f1, f1, f1, f1, f1, f1, f1, ef1, f1}
	err := Run(fSlice, 4, 1)
	if err != nil {
		fmt.Println(err)
	}
}
