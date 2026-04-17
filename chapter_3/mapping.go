package main

import (
	"fmt"
	"maps"
)

func main() {
	//Объявление map
	//var nilMap map[string]int
	//totalWins := map[string]int{}
	//ages := make(map[int][]string, 10)
	comparisonMap()
}

func newMap() {
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)
	fmt.Println(len(teams))
}

func exmaple310() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Linos"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
}

func commaOk() {
	m := map[string]int{
		"hello": 5,
		"world": 6,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ok = m["goodbye"]
	fmt.Println(v, ok)
	m["goodbye"] = 7
	fmt.Println(m["goodbye"])
	//Удаление из map
	fmt.Println(m, len(m))
	delete(m, "world")
	fmt.Println(m, len(m))
	//Очистка m
	clear(m)
	fmt.Println(m, len(m))
}

// Сравнение map
func comparisonMap() {
	m := map[string]int{
		"hello": 5,
		"world": 6,
	}
	n := map[string]int{
		"world": 10,
		"hello": 5,
	}
	fmt.Println(maps.Equal(m, n))
}
