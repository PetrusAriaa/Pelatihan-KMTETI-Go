package main

import "fmt"

func main() {
	// Display text without adding new line
	fmt.Print("Hello, world!")
	fmt.Print("Hello from Golang!")
	fmt.Print("\n")
	fmt.Print("====\n")

	// Add new line with "\n" escape characters
	fmt.Print("Hello, world!\n")
	fmt.Print("Hello from Golang!\n")
	fmt.Print("====\n")

	// Display new line with fmt.Println()
	fmt.Println("Hello, world!")
	fmt.Println("Hello from Golang!")
	fmt.Println("Hello", "world", "from", "Golang!")
}
