//Dynamic type declaration / Type Inference

package main

import "fmt"

func main() {
	var x float64 = 20.0

	y := 45 //dynamic type
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)

}
