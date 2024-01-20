package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go SquareOfNumbers(&wg)

	wg.Wait()
}

func SquareOfNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	ch2 := make(chan int)

	go func() {
		defer close(ch2)

		for val := range RandomNumbers() {
			a := val * val
			ch2 <- a
		}
	}()

	for val := range ch2 {
		fmt.Print(val, ", ")
	}
}

func RandomNumbers() chan int {
	ch1 := make(chan int)

	randNum := rand.Intn(10-5) + 5

	go func() {
		defer close(ch1)
		for i := 0; i < randNum; i++ {
			random := rand.Intn(30-10) + 10
			ch1 <- random
		}
	}()

	return ch1
}