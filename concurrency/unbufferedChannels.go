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

func main() {
	court := make(chan int)
	wg.Add(2)

	go player("Nadal", court)
	go player("Roger", court)

	court <- 1
	wg.Wait()
	fmt.Printf("Game Over.\n")

}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		// only a seed or a different goroutine can unblock below statement
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won!\n", name)
			return
		}
		randomNum := rand.Intn(100)
		if randomNum%13 == 0 {
			fmt.Printf("Player %s missed!\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s hit %v\n", name, ball)
		ball++

		// This is blocking,
		// meaning that below stament wont complete until someone (another goroutine) listens to it on the other end
		court <- ball
	}

}
