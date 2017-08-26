package main

import "fmt"

// scope
var x string = "Hello World"

func main() {
	fmt.Println(x)
	f()
}

func f() {
	fmt.Println(x)
}
