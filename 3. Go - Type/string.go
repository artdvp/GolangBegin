package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello world"[1]) //show byte
	fmt.Println("Hello " + " World")

	var mm string = `1
	2
	3
	4
	5
	6
	`
	fmt.Println(mm)
}
