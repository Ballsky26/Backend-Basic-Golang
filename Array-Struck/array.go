package main

import "fmt"

func main_array() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	fmt.Println(fruits[1])
}
