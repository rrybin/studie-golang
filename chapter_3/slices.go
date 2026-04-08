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
	test()
}

func test() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	//s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	//fmt.Println(slices.Equal(x, s))

}
