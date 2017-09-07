package main

import "fmt"

func main() {
	someOtherName := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(someOtherName))

	/*
		86.6
	*/
}

func average(xs []float64) float64 {
	//panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
