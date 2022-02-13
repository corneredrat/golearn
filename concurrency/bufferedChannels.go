package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	wg sync.WaitGroup
)

const (
	numberOfWorkers = 4
	numberOfTasks   = 20
)

func main() {
	taskManager := make(chan string, numberOfTasks)
	wg.Add(numberOfWorkers)
	for workerID := 0; workerID < numberOfWorkers; workerID++ {
		go worker(workerID, taskManager)
	}
	for task := 0; task < numberOfTasks; task++ {
		taskManager <- fmt.Sprintf("Task-%d", task)
	}
	close(taskManager)
	wg.Wait()
}

func worker(id int, taskManager chan string) {
	defer wg.Done()
	for {
		task, ok := <-taskManager
		if !ok {
			fmt.Printf("No more tasks found, worker %d signing off.\n", id)
			return
		}

		sleepTime := rand.Int63n(10)
		fmt.Printf("Worker %d doing %s for %vs\n", id, task, sleepTime)
		time.Sleep(time.Duration(sleepTime) * time.Second)

		fmt.Printf("Worker %d completed %s\n", id, task)
	}
}
