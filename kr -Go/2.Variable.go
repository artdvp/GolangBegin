package main

import "fmt"

// variable

func main() {
	var myAge int = 20
	myAge2 := 500.1

	var ss bool = true
	txt := "string"
	age1, age2 := 35, 80

	fmt.Println("My Age is ", myAge)
	fmt.Println("My Age2 is ", myAge2)
	fmt.Println("My ss is ", ss)
	fmt.Println("My txt is ", txt)
	fmt.Printf("My age1 age2 is %d , %d", age1, age2)
}
