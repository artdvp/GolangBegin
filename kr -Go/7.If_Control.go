package main

import "fmt"

func main() {
	fmt.Printf("Input Your Number: ")
	var input float64
	fmt.Scanf("%f", &input)

	condition := input > 2
	if condition {
		fmt.Println("input is more than 2")
	} else {
		fmt.Println("input is lesster than 2")
	}

}
