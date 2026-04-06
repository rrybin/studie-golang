package main

import "fmt"

const x int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)
	// x = x + 1
	// y = "bye"
	// fmt.Println(x)
	// fmt.Println(y)
	//Нетипизированная константа
	const i = 10
	var r int = i
	var e float64 = i
	var d byte = i
	fmt.Println(r, e, d)
}
