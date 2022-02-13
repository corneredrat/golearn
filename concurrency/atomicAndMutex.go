package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)

	counter = 0
	go lameWorker(1)
	go lameWorker(2)
	wg.Wait()
	fmt.Printf("Lame Workers are done running. counter Value: %v\n", counter)
	wg.Add(2)

	counter = 0
	go syncedAtomicWorker(1)
	go syncedAtomicWorker(2)
	wg.Wait()
	fmt.Printf("Synced Workers are done running. counter Value: %v\n", counter)

	wg.Add(2)
	counter = 0
	go syncedMutexedWorker(1)
	go syncedMutexedWorker(2)
	wg.Wait()
	fmt.Printf("Mutex enabled Workers are done running. counter Value: %v\n", counter)
}

func lameWorker(id int) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {

		value := counter // if goScheduler moves this goroutine back int the queue at this moment, then we can see race condition
		counter = value + 1

	}
}

func syncedAtomicWorker(id int) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func syncedMutexedWorker(id int) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		value := counter
		counter = value + 1
		mutex.Unlock()
	}
}
