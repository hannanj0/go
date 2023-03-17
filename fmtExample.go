package main

import (
	"fmt"
)

func main() {

	// Using print, scan and sprint from the fmt package in Go
	var name string
	fmt.Println("What is your name?")
	// Saving this using a reference to the name variable
	fmt.Scan(&name)
	// Print hello, with the name input
	fmt.Println("Hello", name)

	var age int
	fmt.Println("What is your age?")
	fmt.Scan(&age)

	// Using printf to print the name and age of the user input
	fmt.Printf("%s is %d years old.\n", name, age)

	// Creating variable newMessage with the Sprintf function

	newMessage := fmt.Sprintf("Hi I am %s and I am %d years old.", name, age)
	fmt.Println(newMessage)

}
