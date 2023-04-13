package main

import "fmt"

func main() {
	var firstName string = "Hello"
	const lastName = "World"

	// Variable bisa Input ulang atau ganti
	firstName = "Aku"
	// Sedangkan Constant tidak bisa Input ulang atau ganti
	// lastName = "Welcome"
	fmt.Print("halo ", firstName, lastName, "!\n") 
}