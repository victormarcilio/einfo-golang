package main

import (
	"fmt"
	"sync"
	"time"
)

var x int

var mut sync.Mutex

func f(a int) {
	for i := 0; i < 50000; i++ {
		mut.Lock()
		fmt.Println(a)
		x = x + 1
		mut.Unlock()
	}
	fmt.Println("terminei", a)
}

func main() {
	fmt.Println("O valor de x eh ", x)
	go f(1)
	go f(2)

	time.Sleep(10 * time.Second)
	fmt.Println("O valor de x eh ", x)
}
