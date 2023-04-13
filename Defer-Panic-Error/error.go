package main

import (
	"fmt"
	"strconv"
)

func main_error() {
	var input string
	fmt.Println("Type some number: ")
	fmt.Scanln(&input)
	var number int
	var err error
	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is input")
		fmt.Println(err.Error())
	}
}
