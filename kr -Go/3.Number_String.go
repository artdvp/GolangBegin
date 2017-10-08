package main

import "fmt"

func main() {
	fnumber := 400
	snumber := 55

	out := fnumber + snumber
	fmt.Println(out)
	fmt.Println(fnumber - snumber)
	fmt.Println(fnumber * snumber)
	fmt.Println(fnumber / snumber)
	fmt.Println(fnumber % snumber)
	fmt.Println(fnumber & snumber)
	fmt.Println(fnumber | snumber)
	fmt.Println(fnumber ^ snumber)
	fmt.Println(fnumber * snumber)
	//fmt.Println(fnumber >> snumber)
	//fmt.Println(fnumber << snumber)
	// string
	p1 := "myDev"
	p2 := " Ops"
	// concatination
	p3 := p1 + p2
	fmt.Println(p3)
	fmt.Println(p3[1:6])
	fmt.Println(p3[0]) // ASCII Code 109
	fmt.Println(p3[4:])
	fmt.Println(p3[:7])
}
