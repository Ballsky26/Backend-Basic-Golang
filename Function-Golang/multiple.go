package main

import (
	"fmt"
	"math"
)

func main_multiple() {
	area, circum := calculated(3.5)
	fmt.Println(area)
	fmt.Println(circum)
}

func calculated(d float64) (float64, float64) { // hitung luas
	var area = math.Pi * math.Pow(d/2, 2) // hitung keliling
	var circumference = math.Pi * d
	// Kembalikan 2 nilai
	return area, circumference
}
