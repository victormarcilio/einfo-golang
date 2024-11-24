package main

import "fmt"

func main() {
	a := 15
	var p *int
	fmt.Println("O pointeiro é nulo?", p == nil)
	p = &a
	fmt.Println(&a) // Imprime o endereço de A
	fmt.Println(p)  // Imprime o valor salvo em P, que é o endereço de A
	fmt.Println(&p) // Imprime endereço de P
	fmt.Println("O pointeiro é nulo?", p == nil)
	fmt.Println("A = ", a)
	*p = 35
	fmt.Println("A = ", a)
	{
		c := 50 // c começa a existir aqui
		fmt.Println(c)
	} // a partir daqui c não é mais acessível
	{
		a := 50 // declaração de uma nova variável com mesmo identificador
		fmt.Println(a)
	} // Depois dessa linha, o identificador 'a' passa a se referir à variável do início da funçao main

	fmt.Println(a)
	fmt.Println("iglobal é ", iGlobal) // escopo de pacote
}
