package main

import "fmt"

func main() {
	i := 1
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println(i)
		}
		i = i + 1
	}

	/*
		for k := 5; k <= 15; k++ {
			fmt.Println(k)
		}
	*/

	/*
		3
		6
		9
		12
		15
		18
		21
		24
		27
		30
		33
		36
		39
		42
		45
		48
		51
		54
		57
		60
		63
		66
		69
		72
		75
		78
		81
		84
		87
		90
		93
		96
		99
	*/
}
