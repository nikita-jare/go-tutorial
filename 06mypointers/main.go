package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	//Why need of pointers?
	//Sometimes when vars passed to function, a copy is created of vars and copy is passed on
	//This created irregularities. To overcome this pointers are used
	//To guarantee actual value is passed on, address of var is passed on

	//* shows data type is pointer
	//var ptr *int
	//default valus is nil
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	//to create pointer on existing variable i.e referencing
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	//operations on pointers changes actual values
	*ptr = *ptr * 2
	fmt.Println("New value is: ", myNumber)
}
