
# Hi Go Begin

> GO TO GO

> ref_1 := [Ebook1](https://www.dropbox.com/sh/is3hwdqa1dpsb99/AADfQKrDju44z2Xx6ukE9WpOa/%E0%B8%AB%E0%B8%81%E0%B8%9A%E0%B8%97%E0%B9%81%E0%B8%A3%E0%B8%81.pdf?dl=0)

## Chap 1 Installation

- https://golang.org/
- https://community.c9.io/t/writing-a-go-app/1725

```
$ sudo rm -rf /opt/go

$ wget https://storage.googleapis.com/golang/goX.X.X.linux-amd64.tar.gz

$ sudo tar -C /opt -xzf goX.X.X.linux-amd64.tar.gz

$ export GOROOT=$opt/

$ export PATH=$PATH:$GOROOT/go/bin
```

## Chap 2 Hello Go

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

## Chap 3. Type

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

## Chap 4 Variable

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

## Chap 5 Control Structures

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

## Chap 6 Arrays, Slices and Map


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

- Make Slices

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

- Function of Slices

	- append

```go
func main(){
	slice1 := []int{1,2,3}
	slice2 := append(slice1,4,5)
	fmt.Println(slice1,slice2)
}
```

	- copy

```
func main(){
	slice3 := []int{1,2,3}
	slice4 := make([]int,2)
	copy(slice3,slice4)
	fmt.Println(slice3,slice4)
}
```
	
