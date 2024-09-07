package main

import "fmt"

func main() {
	s := []int{}
	s = append(s, 1, 2, 3)
	fmt.Println("s:", s)
	fmt.Println("=====")

	f := []string{"mango", "apple", "banana"}
	fmt.Println("f:", f)

	// accessing items by index
	fmt.Println("f[1]", f[1])
	fmt.Println("=====")

	// change value by slice index
	f[1] = "pineaple"
	fmt.Println("f[1]", f[1])
	fmt.Println("=====")

	// add new items into slice
	f = append(f, "durian", "avocado")
	fmt.Println(f)
	fmt.Println("=====")

	// loop through slice items
	for i, item := range f { // range f returns two values, <currentIndex>, <currentItem>
		fmt.Println(item, "| index:", i)
	}
	fmt.Println("=====")

	// delete item from slice
	f2 := []string{}
	for i, item := range f {
		if i != 2 {
			f2 = append(f2, item)
		}
	}
	f = f2
	fmt.Println("f:", f)
}
