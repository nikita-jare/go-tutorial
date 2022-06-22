package main

import "fmt"

//func func_name() return_type {}
//main acts as entry point of program
func main() {
	fmt.Println("Welcome to functions in GoLang")
	greeter()

	//not allowed to write function inside a function
	//func greeterTwo(){
	//	fmt.Println("Another method")
	//}
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 5, 8, 7, 9)
	fmt.Println("Pro Result is: ", proRes)
	fmt.Println("Pro Message is: ", myMessage)
}

//would not execute unless it is called
func greeter() {
	fmt.Println("Namaste from GoLang")
}

func greeterTwo() {
	fmt.Println("Another method")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

//when not sure how many values will be received
func proAdder(values ...int) (int, string) {
	total := 0

	//here values will be a slice
	for _, val := range values {
		total += val
	}

	// can return more than one values
	return total, "Hi Pro result function"
}
