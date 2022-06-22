package main

import "fmt"

//key-value pair, retrieval is very fast

func main() {
	fmt.Println("Maps in GoLang")

	//need non-zero initialised values so make
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	//prints in alphabetical order
	fmt.Println("List of all languages: ", languages)

	fmt.Println("JS shorts for: ", languages["JS"])

	//deleting in maps
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	//loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	//comma ok syntax
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
