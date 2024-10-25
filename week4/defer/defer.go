package main

import "fmt"

func runDefer() {
	fmt.Println("Connecting to database client")
	defer fmt.Println("Connection closed database client")
	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i)
	}

	fmt.Println("Connecting to database cursor")
	defer fmt.Println("Connection to database cursor")

	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i)
	}
}

func main() {
	runDefer()
}
