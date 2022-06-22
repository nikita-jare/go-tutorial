package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Ludo game in SwitchCase in GoLang")

	//generate a seed for random number
	//In Golang, the rand.Seed() function is used to set a seed value to generate pseudo-random numbers.
	//If the same seed value is used in every execution, then the same set of pseudo-random numbers is generated. In order to get a different set of pseudo-random numbers, we need to update the seed value.

	//The seed() method is used to initialize the random number generator.
	//The random number generator needs a number to start with (a seed value), to be able to generate a random number.
	//Use the seed() method to customize the start number of the random number generator.
	//Note: If you use the same seed value twice you will get the same random number twice. See example below

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one and you can open the game")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		//continue to next case too
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots")
	default:
		fmt.Println("What was that?")
	}
}
