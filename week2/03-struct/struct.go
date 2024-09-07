package main

import "fmt"

// declaring type Student with struct
type Student struct {
	Name     string
	Angkatan int16
	IsAsdos  bool
}

func main() {

	// declare new variable with type Student
	var s0 Student
	s0.Name = "Rebecca Brown"
	s0.Angkatan = 2022
	s0.IsAsdos = false

	fmt.Println(s0)
	fmt.Printf("type of s0: %T \n", s0)
	fmt.Println("=====")
	fmt.Println("student0 name:", s0.Name)
	fmt.Println("student0 angkatan:", s0.Angkatan)
	fmt.Println("student0 asdos:", s0.IsAsdos)
	fmt.Println("=====")

	s1 := Student{
		Name:     "John Doe",
		Angkatan: 2021,
		IsAsdos:  true,
	}

	// change value of struct attribute
	s1.IsAsdos = false
	fmt.Println("student1 asdos:", s1.IsAsdos)
}
