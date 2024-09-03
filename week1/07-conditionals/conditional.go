package main

import "fmt"

func main() {
	val1 := 10
	val2 := 15
	val3 := 10.1
	val4 := 10

	// operators:
	// > : is greater than
	// < : is lower than
	// >= : is greater or equal than
	// <= : is lower or equal than
	// == : is equal to
	// != : is not equal to
	fmt.Println(val1 > val2)
	fmt.Println(val3 < float64(val2))
	fmt.Println(val1 == val4)
	fmt.Println(val1 >= val4)
}
