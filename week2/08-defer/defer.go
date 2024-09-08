package main

import (
	"fmt"
	"time"
)

func startCount() {
	defer fmt.Println("Finished [Defer 1]")

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("counter:", i)
	}

	defer fmt.Println("Finished [Defer 2]")

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("[2] counter:", i)
	}
}

func main() {
	startCount()
}
