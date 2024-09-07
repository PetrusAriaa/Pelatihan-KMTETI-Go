package main

import "fmt"

func main() {
	studentAge := map[string]int{
		"Alice": 20,
	}
	studentAge["John"] = 27 // add new key-value pair
	studentAge["Doe"] = 12
	fmt.Println(studentAge)
	fmt.Println("====")

	studentAge["Alice"] = 12 //	change value of key
	fmt.Println(studentAge)
	fmt.Println("====")

	delete(studentAge, "Doe") // delete key from map
	fmt.Println(studentAge)
}
