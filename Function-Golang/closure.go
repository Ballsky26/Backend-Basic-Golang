package main

import "fmt"

func main_closure() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers = []int{2, 3, 2, 2, 3, 4}
	var min, max = getMinMax(numbers)
	fmt.Println(min)
	fmt.Println(max)
}
