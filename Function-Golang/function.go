package main

import (
	"fmt"
	"strings"
)

func main_() {
	var names = []string{"My", "Ballsky"}
	printMessage("Halo", names)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, "")
	fmt.Println(message, nameString)
}
