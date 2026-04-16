package main

import (
	"fmt"
	"slices"
)

func main() {
	var x = []int{10, 20, 30}
	fmt.Println(x)
	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)
	//test()
	//sliceSlices()
	//changeSlices()
	//example37()
	//example38()
	//copySlices()
	//convertArryaToSlices()
	//convertationSlicesToArray()
	//workWithStrings()
	convertationString()
}

func test() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	//s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	x = append(x, 6)
	fmt.Println(len(x), cap(x))
	clear(x)
	fmt.Println(len(x), cap(x))
	fmt.Println(x)
}

func sliceSlices() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}

func changeSlices() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func example37() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "z")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func example38() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func copySlices() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)
	z := make([]int, 2)
	num = copy(z, x)
	fmt.Println(z, num)
	i := make([]int, 2)
	//Копирование последних двух элементов
	copy(i, x[2:])
	fmt.Println(i)
	//Копирование элементов внутри среза
	num = copy(x[:3], x[1:])
	fmt.Println(x, num)
	d := [4]int{5, 6, 7, 8}
	j := make([]int, 2)
	//Копирование массива в срез
	copy(j, d[:])
	fmt.Println(j)
	//Копирование среза в массив
	copy(d[:], x)
	fmt.Println(d)
}

func convertArryaToSlices() {
	xArray := [4]int{4, 5, 6, 7}
	//Перобразование массива в срез
	xSlice := xArray[:]
	fmt.Println(xSlice)
	//Преобразование подмножества
	y := xArray[:2]
	x := xArray[2:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	//Используют общую память
	xArray[0] = 22
	y = xArray[:2]
	x = xArray[2:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

func convertationSlicesToArray() {
	xSlice := []int{1, 2, 3, 4, 5}
	xArray := [5]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
}

func workWithStrings() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println("s[6]:", s[6])
	fmt.Println("b:", b)
	fmt.Println("s[4:7]", s[4:7])
	fmt.Println("s[:5]", s[:5])
	fmt.Println("s[6:]", s[6:])
	//Символы длиной более 1 байта
	s = "Hello ☺"
	fmt.Println(s[4:7])
	fmt.Println(s[:5])
	fmt.Println(s[6:])
	fmt.Println(len(s))
}

func convertationString() {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(a, s, b, s2)
}
