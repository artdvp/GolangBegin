package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for k := 5; k <= 15; k++ {
		fmt.Println(k)
	}

	/*
		1
		2
		3
		4
		5
		6
		7
		8
		9
		10
		5
		6
		7
		8
		9
		10
		11
		12
		13
		14
		15
	*/
}
