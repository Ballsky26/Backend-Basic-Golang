package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main_structs() {
	var s1 student
	s1.name = "Ballsky"
	s1.grade = 2
	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)
}
