package main

import (
	"fmt"
	"sort"
)

//slices are most extensively used than arrays

func main() {
	fmt.Println("Welcome to session on slices")

	//need to initialise in this type of declaration
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	//Type of fruitList is []string

	//Adding to slices
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	//more used syntax
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	//make keyword to define slices
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 12
	highScores[2] = 945
	highScores[3] = 867
	//if try to add 5th value, it will crash as size is 4 in make()

	//using append will realocate memory and all values will be accomodated
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)

	//sorting slices
	sort.Ints(highScores)
	fmt.Println(highScores)

	//boolean operations
	fmt.Println(sort.IntsAreSorted(highScores))

	//remove a value in slices based on indes
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	//[1:3] 1 is inclusive, 3 is not inclusive
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
