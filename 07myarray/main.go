package main

import "fmt"

func main() {

	//arrays not much used in golang
	fmt.Println("Welcome to arrays in golang")

	//declaration, have to give size
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	//not assigned the 2 location, it will print a space
	fmt.Println("Fruit list is: ", fruitList)

	//no matter how many elements, length will always be four
	fmt.Println("Fruit list length is: ", len(fruitList))

	var veggieList = [5]string{"potato", "beans", "mushrooms"}
	fmt.Println("Veggie list is: ", len(veggieList))

	//no sorting, no special operations
}
