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
}
