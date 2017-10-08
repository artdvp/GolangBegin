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

	fmt.Println(add(1, 2, 3))
}

// 6

//func Println(a ...interface{}) (n int,err error)
