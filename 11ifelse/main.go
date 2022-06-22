package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	loginCount := 10
	var result string

	//not allowed to move { onto next line of if statement}
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exaclty 10 login count"
	}

	fmt.Println(result)

	//create var on the go
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	//when assign the value on the go and check condition, ex-web request handling
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {

	//}
}
