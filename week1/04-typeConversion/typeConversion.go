package main

import "fmt"

func main() {
	x := 5
	y := x / 2
	fmt.Println("y:", y) // result is 2, instead of 2.5 because x and y are int
	fmt.Printf("TYPE: x: %T, y: %T \n", x, y)

	xConv := float32(x)
	// y = x / 2							// change value with different types causes error
	yConv := xConv / 2
	fmt.Println("yConv: ", yConv)
	fmt.Printf("TYPE: xConv: %T, yConv: %T \n", xConv, yConv)
}
