package main

import "fmt"

// contoh function tanpa return
func printMessage(msg string) {
	fmt.Println(msg)
}

// contoh function single return
func addNumbers(a int, b int) int {
	res := a + b
	return res
}

// contoh function multiple return
func findIndexAndLength(s []string, name string) (int, string) {
	index := -1
	for i, item := range s {
		if item == name {
			index = i
		}
	}
	return index, name
}

func main() {
	m := "Function dengan Bahasa Golang"
	printMessage(m)
	fmt.Println("=====")

	x := 1
	y := 2
	result := addNumbers(x, y)
	fmt.Println("x + y =", result)
	fmt.Println("=====")

	stud := []string{"Alice", "Bob", "Chris", "Dwayne"}
	idx, n := findIndexAndLength(stud, "Chris")
	fmt.Println("Index of student", n, ":", idx)
}
