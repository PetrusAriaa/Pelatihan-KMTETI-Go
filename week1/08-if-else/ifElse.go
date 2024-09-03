package main

import "fmt"

func main() {
	s := 88
	if s >= 90 {
		fmt.Println("Grade: A")
	} else if s >= 80 {
		fmt.Println("Grade: B")
	} else if s >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}
}
