package main

import "fmt"

type Cor uint8

const (
	Amarelo Cor = iota
	Azul
	Branco
	Lilas
	Rosa
)

const PI = 3.1415872371929817309123810231271379
const X = 23323

type Tamanho uint8

const (
	PP = 1 // Todos os valores serão inicializados iguais a esse se a atribuição for omitida
	P
	M
	G
	GG
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius {
	return Celsius((f + 32) * 5 / 9)
}

func main() {
	var tempCelsius Celsius = 30.5
	var tempFah Fahrenheit = 86.9

	fmt.Println(tempCelsius == Celsius(tempFah))

	fmt.Println("Cores:", Amarelo, Azul, Rosa)
	fmt.Println("Tamanhos: ", PP, P, M, G)
}
