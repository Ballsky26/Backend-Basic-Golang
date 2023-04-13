package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate2(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Kata tidak boleh Kosong")
	}
	return true, nil
}
func main_() {
	var name string
	fmt.Println("Type some name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("Halo", name)
	} else {
		fmt.Println(err.Error())
	}
}
