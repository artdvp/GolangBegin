package main

import "fmt"

var x int = 5

func main() {

	fmt.Println(f1())
}

func f1() int {
	return f2()
}

func f2() int {
	return 1
}

func f3() (r int) {
	r = 1
	return
}
