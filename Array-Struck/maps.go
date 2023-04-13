package main

import "fmt"

func main_maps() {
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40 // menggunakan Key
	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"]) // apabila kita memanggil nilai yg tidak ada maka hasilnya 0
}
