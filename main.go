package main

import "fmt"

// variáveis globais são variáveis do pacote e o compilador não vai reclamar caso não as utilize abaixo
var fGlobal = 3.5

// A inicialização abaixo não é permitida para variáveis globais
//bGlobal := true

func main() {
	fmt.Println("Olá mundo")

	var a, a2 int // a é inicializado com o valor 'zero' do tipo int
	var b bool    // diferentemente de outras linguagens, no Go o(s) nome(s) da(s) variável(is) vem antes e o tipo vem depois
	var s string
	fmt.Println(a, a2, b, s)
	fmt.Printf("%v %v '%v'\n", a, b, s)

	var c = 10
	var s2 = "alguma string"
	fmt.Println(c, s2)

	// Declara e inicializa ao mesmo tempo
	s3 := "terceira string"

	fmt.Println(s3)

	// A linha abaixo funciona porque 'z' está sendo declarada. s2 é reutilizada
	z, s2 := 5, "segunda string"

	// A linha comentada abaixo não compila porque não existe nenhuma variável nova a esquerda do :=
	// z, s2 := 5, "segunda string"
	fmt.Println(z)
}
