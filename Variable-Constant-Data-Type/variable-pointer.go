package main

import "fmt"

func main() {
var numberA int = 4
var numberB *int = &numberA
// Jika di panggil menggunakan variable pointer maka variable sebelumnya berubah
*numberB = 8
// Jika ingin memanggil nilainya menggunakan tanda *
fmt.Print("halo ", numberA, "!\n") 
}