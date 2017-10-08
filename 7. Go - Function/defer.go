package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second() // execute after first function complete
	first()
}

/*
1st
2nd
*/
