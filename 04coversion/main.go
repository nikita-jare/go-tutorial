package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to ou Pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	//read till \n
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	//As i/p of type string, it cannot add 1 to it
	//numRating = input + 1
	//numRating, err := strconv.ParseFloat(input, 64)
	// as passing only input to strconv will through error becaus input will include \n, we need to trim it
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	//error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
