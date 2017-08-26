package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 33
	y[1] = 44
	y[2] = 66
	y[3] = 99
	y[4] = 111

	var total float64 = 0
	var total2 float64 = 0
	var total3 float64 = 0
	var total4 float64 = 0
	//for i := 0; i < 5; i++ {
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	//fmt.Println(total / 5)
	fmt.Println(total / float64(len(y)))

	for _, value := range y {
		total2 += value
	}
	fmt.Println(total2 / float64(len(y)))

	// defind Array
	q := [5]float64{98, 44, 35, 67, 22}
	for _, value := range q {
		total3 += value
	}
	fmt.Println(total3 / float64(len(q)))

	z := [5]float64{
		98,
		44,
		35,
		67,
		//22, // for comment
	}
	for _, value := range z {
		total4 += value
	}
	fmt.Println(total4 / float64(len(z)))
}
