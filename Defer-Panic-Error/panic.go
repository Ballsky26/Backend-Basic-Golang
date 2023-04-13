package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate1(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Kata tidak boleh Kosong")
	}
	return true, nil
}

func main__() {
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
