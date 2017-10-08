
# Hi Go Begin

> GO TO GO

> ref_1 := [chap 1-6](https://www.dropbox.com/sh/is3hwdqa1dpsb99/AADfQKrDju44z2Xx6ukE9WpOa/%E0%B8%AB%E0%B8%81%E0%B8%9A%E0%B8%97%E0%B9%81%E0%B8%A3%E0%B8%81.pdf?dl=0)

> ref_2 := [chap 7](https://www.dropbox.com/sh/is3hwdqa1dpsb99/AAC78qWzQ4wbPYBMTAPBcf3Xa/%E0%B8%9A%E0%B8%97%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%80%E0%B8%88%E0%B9%87%E0%B8%94.pdf?dl=0)

> ref_3 := [chap 8](https://www.dropbox.com/sh/is3hwdqa1dpsb99/AAAK_CtDN8qeKf_gwc5lR4dga/%E0%B8%9A%E0%B8%97%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%81%E0%B8%9B%E0%B8%94.pdf?dl=0)

> ref_4 := [chap 9](https://www.dropbox.com/sh/is3hwdqa1dpsb99/AADkVBUqqvVrqcTi1lEAj2sla/%E0%B8%9A%E0%B8%97%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%80%E0%B8%81%E0%B9%89%E0%B8%B2.pdf?dl=0)
## Chapter 1 - Installation

- https://golang.org/
- https://community.c9.io/t/writing-a-go-app/1725

```
$ sudo rm -rf /opt/go

$ wget https://storage.googleapis.com/golang/goX.X.X.linux-amd64.tar.gz

$ sudo tar -C /opt -xzf goX.X.X.linux-amd64.tar.gz

$ export GOROOT=$opt/

$ export PATH=$PATH:$GOROOT/go/bin
```

## Chapter 2 - Hello Go

```go
package main

import "fmt" //format

func main() {
	fmt.Printf("Hello, the Golang\n") 
}
```

```sh
$ go run hello.go

$ go build hello.go

$ ./hello

$ go help
```

## Chapter 3 - Type

### Numbers

- Integer
	- uint8 (unsigned integer -- positive number + 0)
    - uint16
    - uint32
    - uint64
    - int8 (signed integer -- nagative 0 positive number)
    - int16
    - int32
    - int64
    - byte = uint8
    - rune = int32
    - for architech uint,int uintptr
        
    - Floating-point
        - float32
        - float64
        - complex64 --> (imaginary number)
        - complex128
        
### String

- string "ABC"

### Boolean
	
- true 
- false

## Chapter 4 - Variable

### Constant
    
```go
package main

import "fmt"

func main() {
	const x string = "Hello World"
	//x = "Some thing" can not asign constant
	fmt.Println(x)

	// Difining Multiple Variable
	/*var (
		a = 4
		b = 100
		c = 6
	)
	*/
	fmt.Println("Enter a Number :")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

```

## Chapter 5 - Control Structures

### For loop
    
```go
package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for k := 5; k <= 15; k++ {
		fmt.Println(k)
	}
}

```

### IF / IF ELSE

```go
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

```

## Chapter 6 - Arrays, Slices and Map


### Arrays

```go
var x [5]int
```

```go
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
```

### Slices

> Make Slices

```go
var x []float64

x := make([]float64,5)

x := make([]float64,5,10) => 10 is length of array at slice reference

[0][1][2][3][4][5][6][7][8][9]
[-------xx-------]

> make slice with [low:high]

arr := [5]float64{1,2,3,4,5}

x := arr[0:5] // [1,2,3,4,5]

x := arr[1:4] // [2,3,4]

// not defind low
x := arr[0:] => arr[0:len(arr)] 

x := arr[:5] => arr[0:5] 

x := arr[:] => arr[0:len(arr)]

```

#### Function of Slices

> append

```go
func main(){
	slice1 := []int{1,2,3}
	slice2 := append(slice1,4,5)
	fmt.Println(slice1,slice2)
}
```

> copy

```go
func main(){
	slice3 := []int{1,2,3}
	slice4 := make([]int,2)
	copy(slice3,slice4)
	fmt.Println(slice3,slice4)
}
```

> map

```go
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

	//delete(x,1)
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
```
	
## Chapter 7 - Function

```go
package main

import "fmt"

func main() {
	someOtherName := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(someOtherName))

	/*
		86.6
	*/
}

func average(xs []float64) float64 {
	//panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

```

```go
package main

import "fmt"

var x int = 5

func main() {

	f()
}

func f() {
	fmt.Println(x)
}
```

```go
package main

import "fmt"

var x int = 5

func main() {

	fmt.Println(f1())
}

func f1() int {
	return f2()
}

func f2() int {
	return 1
}

func f3() (r int) {
	r = 1
	return
}
```

```go
package main

import "fmt"

func f1(int, int) {
	return 5, 6
}

func main() {

	x, y := f()
}
```

> Variadic Function : multi parameter

```go
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

```


```go
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

```

> Closure : function unnecessary have name can be assign to variable

```go
package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 5))
}

// 6
```

```go
package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	/*
		1
		2
	*/
}
```

```go
package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

```

> Recursive function


```go
package main

import (
	"fmt"
)

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {

	var k uint
	k = 10
	fmt.Println(factorial(k))
}

```

> Defer , Panic & Recover

defer execute afeter first function complete

```go
package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second() // execute after first function complete
	first()
}
```

> Panic & Recover

```go
package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")

}
```

## Chapter 8 - Pointers

```go
package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}

// 0

```

> new

```go
package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}

// 1

```

-- chap 9 ...