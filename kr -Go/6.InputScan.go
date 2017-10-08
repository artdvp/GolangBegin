package main

import "fmt"

func main() {
	fmt.Printf("Input Your Number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("input is :", output)
}
