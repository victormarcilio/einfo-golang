package main

import "fmt"

func main() {
	var array [5]int = [5]int{3, 5, 10} // os ultimos 2 elementos serão 'zero'
	array2 := [...]int{3, 5, 10}
	fmt.Println(array, array2)

	ar2 := [2]int{1, 5}
	ar3 := [2]int{1, 5}
	ar4 := [3]int{1, 5, 0}
	ar5 := [2]int{1, 3}
	ar4 = ar4 // não faça isso! apenas pra que o compilador deixe passar
	fmt.Println(ar2 == ar3)
	fmt.Println(ar2 == ar5)
	ar5[1] = 5
	fmt.Println(ar2 == ar5)
	// O tamanho do array faz parte do tipo e não é possível fazer comparações entre tipos diferentes
	//fmt.Println(ar3 == ar4)

	x := make([]int, 10) // cria slice com 10 elementos inicializados com o valor 0
	fmt.Println("x: ", x, len(x), cap(x))
	y := make([]int, 2, 10) // cria slice com 2 elementos inicializados com 0 e capacidade total para 10
	fmt.Println("y: ", y, len(y), cap(y))
}
