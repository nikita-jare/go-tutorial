package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Saturday"}

	fmt.Println(days)

	//no ++d, only d++
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//i returns index, in other languages it return values
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index is %v and values is %v\n", index, day)
	}

	rougueValue := 1
	//similar to while
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			//break
			continue
		}

		fmt.Println("Values is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at goto")
}
