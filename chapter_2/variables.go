package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	//Bool
	var flag bool
	var isAwesome = true
	fmt.Println(flag, "\n", isAwesome)
	var x int = 10
	x *= 2
	fmt.Println(x)
	cmx()
	str()
	conversion()
	conversionTwo()
	declarationVriables()
}

// Complex digital
func cmx() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
}

// String and rune
func str() {
	var myFirstInitial rune = 'J'
	var myLastInitial int32 = 'J'
	fmt.Println(myFirstInitial, " ", myLastInitial)
}

// Conversion type
func conversion() {
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println("sum1 = ", sum1, "\nsum2 = ", sum2)
}

func conversionTwo() {
	var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b
	fmt.Println(sum3, " ", sum4)
}

func declarationVriables() {
	var x = 10
	var z, y int = 20, 30
	var i, j = 40, "hello"
	fmt.Println(x, z, y, i, j)
}
