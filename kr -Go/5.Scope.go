// Scope
// Global
package main

import "fmt"

var gVal int = 500

func main() {
	iVal := 40
	fmt.Println("Global -", gVal) // Global
	fmt.Println("in var -", iVal)
	oh(iVal)
}

func oh(inx int) {

	fmt.Println("Val")
	fmt.Println("in var -", inx)
}
