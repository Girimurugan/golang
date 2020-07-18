package main

import (
	"fmt"
	"sync"
)

func lowercase() {
	fmt.Println("Lower case")
}

func uppercase() {
	fmt.Println("Upper case")
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting the Exec case")

	go func() {
		lowercase()
		wg.Done()
	}()

	go func() {
		uppercase()
		wg.Done()
	}()

	fmt.Println("Waiting")
	wg.Wait()

	fmt.Println("Finished")

}
