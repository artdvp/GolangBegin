package main

import "fmt"

func main() {
	// var x map[string]int // x is map string of Numuric
	// x["key"] = 10
	// fmt.Println(x)
	/*
			panic: runtime error: assignment to entry in nil map

		    goroutine 1 [running]:
		    runtime.panic(0x4999e0, 0x56579d)
		        /usr/lib/go/src/pkg/runtime/panic.c:266 +0xb6
		    main.main()
		        /home/ubuntu/workspace/GolangBegin/6. Go - Arrays-Slices-Map/map.go:7 +0x53
		    exit status 2
	*/
	//error// map must assign begin value before working

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	/*
	   10
	*/

	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1])
	/*
	   10
	*/

	//delete
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	/*
		Lithium
	*/
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	/*
		false
	*/

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	/*
		false
	*/

	elements2 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements2["He"])
	/*
		Helium
	*/

	elements3 := map[string]map[string]string{

		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
	}
	fmt.Println(elements3["Li"])

	/*
		map[name:Lithium state:solid]
	*/
}
