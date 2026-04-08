package main

import "fmt"

func main() {
	var x [3]int
	fmt.Println(x)
	var y = [3]int{10, 20, 30}
	fmt.Println(y)
	var z = [12]int{1, 5: 4, 6, 10: 100, 25}
	fmt.Println(z)
	var j = [...]int{1, 2, 3}
	fmt.Println(j)
	var i [2][3]int
	i[0][0] = 123
	fmt.Println(i)
}
