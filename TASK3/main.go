package main

import (
	"fmt"
	"os"
	"sync"
)

func writeToFile(filename string, numbers []int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, num := range numbers {
		fmt.Fprintf(file, "%d ", num)
	}
}

func main() {
	var wg sync.WaitGroup

	
	numbersFile1 := make([]int, 0)
	for i := 0; i <= 100; i += 5 {
		numbersFile1 = append(numbersFile1, i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		writeToFile("file1.txt", numbersFile1)
	}()

	numbersFile2 := make([]int, 0)
	for i := 0; i <= 100; i += 3 {
		numbersFile2 = append(numbersFile2, i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		writeToFile("file2.txt", numbersFile2)
	}()

	wg.Wait()

	numbersFile3 := make([]int, 0)
	for _, num := range numbersFile1 {
		if num%15 == 0 {
			numbersFile3 = append(numbersFile3, num)
		}
	}
	for _, num := range numbersFile2 {
		if num%15 == 0 {
			numbersFile3 = append(numbersFile3, num)
		}
	}

	writeToFile("file3.txt", numbersFile3)
}
