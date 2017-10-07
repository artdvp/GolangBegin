// Expression Switch
package main

import "fmt"

func main() {
	/* local variable defintion */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Println("Well done\n")
	case grade == "D":
		fmt.Println("You passed\n")
	case grade == "F":
		fmt.Println("Better try again\n")
	default:
		fmt.Println("Invalid grade\n")
	}
	fmt.Printf("Your grade is %s\n", grade)
}
