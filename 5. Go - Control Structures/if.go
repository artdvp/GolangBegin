package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {

		if i%2 == 0 {
			fmt.Println(i, "Odd")
		} else {
			fmt.Println(i, "Even")
		}

		i = i + 1
	}

	for k := 5; k <= 15; k++ {

		if k%2 == 0 {
			fmt.Println(k, "devisible by 2")
		} else if k%3 == 0 {
			fmt.Println(k, "devisible by 3")
		}

	}

	/*
		1 Even
		2 Odd
		3 Even
		4 Odd
		5 Even
		6 Odd
		7 Even
		8 Odd
		9 Even
		10 Odd
		6 devisible by 2
		8 devisible by 2
		9 devisible by 3
		10 devisible by 2
		12 devisible by 2
		14 devisible by 2
		15 devisible by 3
	*/
}
