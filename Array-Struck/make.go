package main

import "fmt"

func main() {
	var fruits = make([]string, 3)
	fruits[0] = "Apple"
	fruits[1] = "Melon"
	fruits[2] = "Grappe"
	fmt.Println(fruits)
}
