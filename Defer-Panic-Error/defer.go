package main

import "fmt"

func main_defer() {
	defer fmt.Println("Halo") // Akan di baca dibawah
	fmt.Println("Selamat Datang")
}
