package main

import (
	"fmt"
	"sync"
)

// "fmt"

func printPrime(wg *sync.WaitGroup, name string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%v : %v\n", name, outer)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Printf("Creating goroutines\n")
	go printPrime(&wg, "A")
	go printPrime(&wg, "B")
	wg.Wait()
	fmt.Printf("Terminating program...\n")
}
