package main

import (
	"errors"
	"fmt"
	"strings"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Kata tidak boleh Kosong")
	}
	return true, nil
}

func main() {
	defer catch()
	var name string
	fmt.Println("Type some Name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("Halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
