package main

import (
	"fmt"
	"maps"
)

func main() {
	var m map[int]string

	// A linha abaixo vai causar um panic porque o mapa é nulo
	//m[45] = "quebrou"
	fmt.Println(m == nil)
	m = make(map[int]string)
	fmt.Println(m == nil)

	m2 := map[int]string{
		1: "um",
		2: "dois",
		3: "três",
	}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "um",
		2: "dois",
		3: "três",
	}
	fmt.Println(maps.Equal(m2, m3))
	delete(m3, 2)
	fmt.Println(maps.Equal(m2, m3))
	fmt.Println(m2, m3)

	value, ok := m3[1]
	fmt.Println(value, ok)
}
