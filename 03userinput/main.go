package main

import (
	"bufio"
	"fmt"
	"os"
)

//bufio package - buffer which can read from input/output
//os package
func main() {
	welcome := "Hello, welcome to User Input"
	fmt.Println(welcome)

	//walrus operator when not sure of what type of input will be received
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err ok syntax
	// no try catch in golang, you either get input or receive an error and hold input in var

	//ReadString - read till \n
	//If do not care about error
	input, _ := reader.ReadString('\n')

	//If care about error, if everything is Okay, value will be in input, if error, stored in err
	//input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	//Gives 5 input as string
	fmt.Printf("Type of this rating is %T", input)
}
