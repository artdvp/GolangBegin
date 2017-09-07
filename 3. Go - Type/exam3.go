package main

import "fmt"

func main() {
	// number is the most of binary numeral system -> 8 position (hint : 101-1 = 9,102-1 = 99)
	// 11111111
	fmt.Println("11111111 =>", 1+2+4+8+16+32+64+128)
	fmt.Println("32132 x 42452 = ", 32132*42452)
	fmt.Println(len("Hey Hello Go na ja jub jub"))
	fmt.Println((true && false) || (false && true) || !(false && false))

	/*
			 11111111 => 255
			 32132 x 42452 =  1364067664
			 26
		 	 true
	*/
}
