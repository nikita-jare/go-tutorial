//structs are alternative versions of classes
package main

import "fmt"

func main() {
	fmt.Println("Session on structs in golang")

	//no inheritance, no super or parent-child classes in golang
	nikita := User{"Nikita", "Nikita@go.dev", true, 24}
	fmt.Println(nikita)

	//gives key-value pairs with %+v
	fmt.Printf("Nikita's detals are: %+v\n", nikita)

	fmt.Printf("Name is %v and Email is %v.", nikita.Name, nikita.Email)
}

//Name of struct first letter capital
//even vars in struct as they will be public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
