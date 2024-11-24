package main

import (
	"fmt"
	"log"
)

type Operacao func(int, int) int

type FnUnaria func(int) int

type cobaia struct{}

func (c *cobaia) Error() string { return "tentativa de divisao por zero" }

func realizaOperacao(op Operacao, op1, op2 int) int {
	if op == nil {
		return -1
	}
	return op(op1, op2)
}

func quadrado(x int) int {
	return x * x
}

func soma(x, y int) int {
	return x + y
}

func subtracao(x, y int) int {
	return x - y
}

func multiplicacao(x, y int) int {
	return x * y
}

func divisao(x, y int) int {
	return x / y
}

func divisaoEmodulo(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, &cobaia{}
	}
	return x / y, x % y, nil
}

// Uma função que recebe um inteiro e retorna uma nova função de um único parametro que deve calcular a multiplicação de ambos os parâmetros
func criaMultiplicador(multiplicador int) FnUnaria {
	return func(x int) int {
		return multiplicador * x
	}
}

func max(a int, restante ...int) int {
	maior := a
	for _, v := range restante {
		if v > maior {
			maior = v
		}
	}

	return maior
}

func incV(x int) {
	x++
}

func incP(x *int) {
	*x++
}

func main() {

	a := 5
	fmt.Println(a)
	incV(a)
	fmt.Println(a)
	incP(&a)
	fmt.Println(a)

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recuperei do erro %+v", err)
		} else {
			fmt.Println("tudo em ordem!")
		}
	}()

	defer func() {
		fmt.Println("Print do segundo defer")
	}()

	defer func() {
		fmt.Println("Print do terceiro defer")
	}()

	dobra := criaMultiplicador(2)
	triplica := criaMultiplicador(3)
	fmt.Println(dobra(5), triplica(10))

	f := func() {
		res := dobra(5)
		fmt.Println(res)
	}
	f()
	m := make(map[string]Operacao)
	m["*"] = multiplicacao
	m["+"] = soma
	m["-"] = subtracao
	m["/"] = divisao
	A := 5
	B := 3
	fmt.Println(quadrado(A))
	var operacao string
	fmt.Println("Qual operacao deseja realizar? (+, -, / ou *)")
	fmt.Scan(&operacao)
	res := realizaOperacao(m[operacao], A, B)

	fmt.Println("O resultado de ", A, operacao, B, "deu ", res)

	op := subtracao
	fmt.Println("A subtracao eh", op(A, B))
	op = soma
	fmt.Println("A soma eh", op(A, B))
	op = multiplicacao
	fmt.Println("A multiplicacao eh", op(A, B))
	op = divisao
	fmt.Println("A divisao eh", op(A, B))

	div, mod, err := divisaoEmodulo(30, 0)
	if err != nil {
		log.Printf("divisao por zero: %+v", err)
	}

	div, mod, _ = divisaoEmodulo(30, 0)

	fmt.Println("div = ", div, " e mod = ", mod)

	fmt.Println("O maior eh ", max(5))
	fmt.Println("O maior eh ", max(3, 5))
	fmt.Println("O maior eh ", max(200, 3, 5))
}
