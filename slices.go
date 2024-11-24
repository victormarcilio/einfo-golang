package main

import (
	"fmt"
	"slices"
)

func main() {
	var c []int
	a := []int{1, 3, 5}
	b := []int{1, 3, 5}
	fmt.Println(a == nil, b == nil, c == nil)
	fmt.Println(slices.Equal(a, b))
	fmt.Println(len(b), cap(b))
	c = append(b, 10, 20, 30)
	fmt.Println("b: ", len(b), cap(b))
	fmt.Println("c: ", len(c), cap(c))
	fmt.Println(b, c)
	b[1] = 6
	fmt.Println(b, c)
	fmt.Println(c[2:4]) // da posição 2 até 3
	fmt.Println(c[2:])  // de 2 até o final
	fmt.Println(c[:4])  // até posição 3
}
