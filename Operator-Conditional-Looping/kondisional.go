package main

import "fmt"

func main_kondisional() {
	var point = 2 
	if(point == 10) {
		fmt.Println("Lulus dengan Nilai A")
	}else if (point > 5){
		fmt.Println("Lulus dengan Nilai B")
	}else if(point == 4) {
		fmt.Println("Hampir Lulus")
	}else {
		fmt.Println("Tidak Lulus, Nilai anda adalah", point)
	}
}