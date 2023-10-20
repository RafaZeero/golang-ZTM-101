package main

import (
	"fmt"
	"math/rand"
)

func generateRandomInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()

	return out
}

// generate a finite amount of random numbers
func generateRandomIntn(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()

	return out
}

func main() {
	randInt := generateRandomInt(1, 100)

	fmt.Println("generate random int infinite")

	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	randIntnRange := generateRandomIntn(2, 1, 100)

	fmt.Println("generate random int finite")

	for i := range randIntnRange {
		fmt.Println(i)
	}

	randIntn := generateRandomIntn(3, 1, 100)

	for {
		// select {
		// case i, ok := <-randIntn:
		i, ok := <-randIntn
		if !ok {
			fmt.Println("channel closed")
			return
		}
		fmt.Println(i)
		// }
	}
}
