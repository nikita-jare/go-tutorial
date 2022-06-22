package main

import "fmt"

func main() {
	//moment you write defer keyword, the line 6 will get cut out and execute at the end of function
	defer fmt.Println("World1")
	defer fmt.Println("World2")
	defer fmt.Println("World3")
	fmt.Println("Hello")
	defer fmt.Println("World")

	//defer works in LIFO
	//World1, World2, World3, World
	//LIFO: World, World3, World2, World1

	myDefer()
}

//0, 1, 2, 3, 4
//hello, 43210, World, World3, World2, World1
//check again
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
