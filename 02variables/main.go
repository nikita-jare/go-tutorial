package main

import "fmt"

//inside any method, allowed to use walarus operator, outside not allowed
//var jwtToken int = 30000

//constant cannot be changed
//First character capital of Public variable
const LoginToken string = "cdbuhbcuhdj"

func main() {
	var username string = "nikita"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is od type: %T \n", isLoggedIn)

	//unint8 contains conly 0 to 255
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45546673
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//If no value is initialised, bydefault it has zero value
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit way of declaring vars
	//lexer assigns datatype according to the value assigned, cannot change later
	var website = "learncode"
	fmt.Println(website)

	//Walrus operator
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	//Printing Public variable
	fmt.Println(LoginToken)
	fmt.Printf("Variable Type: %T \n", LoginToken)
}
