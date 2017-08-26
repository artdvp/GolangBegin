package main

import "fmt"

func main() {

	fmt.Println("---- Fahrenheit to Celsius ----")
	fmt.Println("Enter number F : ")

	var input float64
	fmt.Scanf("%f", &input)

	//output := ((input - 32) * (5 / 9))
	output := (input - 32) * 5 / 9

	fmt.Println(output)
}
