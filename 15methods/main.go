package main

//when functions go in classes/structs they are called methods
//structs are alternative versions of classes

import "fmt"

func main() {
	fmt.Println("Session on structs in golang")

	//no inheritance, no super or parent-child classes in golang
	nikita := User{"Nikita", "Nikita@go.dev", true, 24}
	fmt.Println(nikita)

	//gives key-value pairs with %+v
	fmt.Printf("Nikita's detals are: %+v\n", nikita)

	fmt.Printf("Name is %v and Email is %v.", nikita.Name, nikita.Email)
	nikita.GetStatus()
	nikita.NewMail()
	//does not change email value globally
	fmt.Printf("Nikita's detals are: %+v\n", nikita)

}

//Name of struct first letter capital
//even vars in struct as they will be public
//if first letter not capital, the var becomes private
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//method
func (u User) GetStatus() {
	fmt.Println("Is use active: ", u.Status)
}

//passes on a copy of object, so it changes locally
func (u User) NewMail() {
	u.Email = "nikita-jare@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
