package main

import "fmt"

type sliceType []int

func main() {
	slice := sliceType{}

	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range ints {
		slice = append(slice, ints[i])
	}

	for i := range ints {
		if slice[i]%2 == 0 {
			fmt.Println(slice[i], "is odd")
		} else {
			fmt.Println(slice[i], "is even")
		}
	}
}
