package main

import "fmt"

func main() {
	// for Loop
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, index:", i)
	}
	fmt.Println("\n====")

	// while-like loop
	counter := 0
	for {
		if counter == 4 { // stop loop when counter reaches 4
			break
		}
		fmt.Println("counter: ", counter)
		counter++
	}
	fmt.Println("\n====")

	// using continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("current: ", i)
	}
}
