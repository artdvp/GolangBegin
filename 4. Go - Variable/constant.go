package main

import "fmt"

func main() {
	const x string = "Hello World"
	//x = "Some thing" can not asign constant
	fmt.Println(x)

	// Difining Multiple Variable
	/*var (
		a = 4
		b = 100
		c = 6
	)
	*/
	fmt.Println("Enter a Number :")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

	/*
	 Hello World
	 Enter a Number :
	 6
	 12
	*/
}
