package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {

	xs := []int{1, 2, 3}
	fmt.Println(add(xs...)) //send multi parameter with slice x...
}

// 6
