package main

import "fmt"

func main() {

	fmt.Println("---- Meter to Foot ----")
	fmt.Println("Enter number F : ")
	const foot float32 = 0.3048
	var input float32
	fmt.Scanf("%f", &input)

	output := input * foot

	fmt.Println(output, " m")

	/*
		---- Meter to Foot ----
		Enter number F :
		78
		23.7744  m
	*/
}
