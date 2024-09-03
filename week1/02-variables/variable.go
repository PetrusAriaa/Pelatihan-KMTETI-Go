package main

import "fmt"

func main() {
	var fruit = "Avocado" // implicit typing
	var count int8 = 20   // explicit type
	var sold int16        // empty variable declaration

	fmt.Println(fruit)
	fmt.Println(count)
	fmt.Println(sold)
	fmt.Println("====")

	newFruit := "Waterlemon" // using := operator
	newFruitCount := 100
	count = 20 // using := operator to declared variable causes error
	fmt.Println(newFruit)
	fmt.Println(newFruitCount)
	fmt.Println(count)
	// Only use := for declaration, use = to change value
}
