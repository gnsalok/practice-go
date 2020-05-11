package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//unfurling the slice
	sum := sum(xi...)
	fmt.Println("The total is : ", sum)

}

func sum(xi ...int) int {
	fmt.Printf("Type of xi %T\n", xi)
	total := 0

	for _, v := range xi {
		total += v
	}
	return total

}
