package main

import "fmt"

func main() {
	x := 15
	y := 200

	switch {
	case y > x:
		fmt.Println("Y eh maior que X")
		fallthrough
	case x < 15:
		fmt.Println("X eh menor que 15")
	}

}
