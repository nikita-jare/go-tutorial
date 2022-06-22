//io package, defer function to close files

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - GoTutorial"
	file, err := os.Create("./mygotutorial.txt")

	// if err != nil {
	// 	//panic will shutdown program and throw error
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./mygotutorial.txt")
}

func readFile(filename string) {
	//creation is os operation
	//read and write has ioutil
	//data will come in bytes format
	databyte, err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	//this will print data in bytes
	//fmt.Println("Text data inside file is:\n", databyte)

	fmt.Println("Text data inside file is:\n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
