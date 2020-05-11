package main

import (
	"fmt"
)

// function to return sum of float array value.
func sumArr(arr *[5]float64) float64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

func main() {

	arr := [5]float64{1, 2, 3, 4, 5}
	result := sumArr(&arr)
	fmt.Println(result)
}
