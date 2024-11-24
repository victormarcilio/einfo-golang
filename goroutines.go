package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func contar(prefix string) {
	for i := 0; i < 200; i++ {
		fmt.Println(prefix, i)
	}
	wg.Done()
}

func main() {

	wg.Add(2)
	go contar("c1")
	go contar("c2")

	fmt.Println("esperando so caras...")
	/*


	 */
	wg.Wait()
	fmt.Println("os caras terminaram")
}
