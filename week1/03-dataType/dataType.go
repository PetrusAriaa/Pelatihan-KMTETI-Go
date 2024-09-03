package main

import "fmt"

func main() {
	newInteger := 8
	nameString := "Fufu"
	newBoolean := true
	fmt.Println("newInteger:", newInteger)

	// Using fmt.Printf to display formatted string
	fmt.Printf("newInteger: %d, type: %T \n", newInteger, newInteger)
	fmt.Printf("nameString: %s, type: %T \n", nameString, nameString)
	fmt.Printf("newBoolean: %v, type: %T \n", newBoolean, newBoolean)

	// variable naming in golang
	// semantic naming is not necessary for local variables
	// use understandable naming for exported variables (later)
	d := "Monday"
	dt := 1
	m := 12
	Y := 2024
	fmt.Printf("%s, %d-%d-%d \n", d, dt, m, Y)

	// default values
	var emptyInt8 uint8
	var otherEmptyString string
	var emptyFloat64 float32
	var emptyBool bool

	fmt.Printf("emptyInt8: %d, otherEmptyString: %s, emptyFloat64: %0.2f, emptyBool: %v \n",
		emptyInt8, otherEmptyString, emptyFloat64, emptyBool)
}
