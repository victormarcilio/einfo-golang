package main

import (
	"fmt"
	"time"
)

type Pedido struct {
	tempoPreparo time.Duration
	tempoEntrega time.Duration
}

type Pizza struct {
	tempoEntrega time.Duration
}

var pedidosPendentes = make(chan Pedido, 3)

var pizzasProntas = make(chan Pizza, 3)

func cozinheiro(nome string) {
	for {
		pedido := <-pedidosPendentes
		fmt.Printf("%s preparando %+v\n", nome, pedido)
		time.Sleep(pedido.tempoPreparo * time.Second)
		fmt.Printf("%s terminou %+v\n", nome, pedido)
		pizzasProntas <- Pizza{pedido.tempoEntrega}
	}
}

func motoboy(nome string) {
	for {
		pizza := <-pizzasProntas
		fmt.Printf("%s saiu para entregar %+v\n", nome, pizza)
		time.Sleep(pizza.tempoEntrega * time.Second)
		fmt.Printf("%s retornou da entrega de %+v\n", nome, pizza)
	}
}

func main() {
	var op int
	for {
		fmt.Println("O que deseja fazer?")
		fmt.Println("1 - Registrar pedido")
		fmt.Println("2 - Registrar cozinheiro")
		fmt.Println("3 - Registrar motoboy")
		fmt.Scan(&op)
		switch op {
		case 1:
			var preparo, entrega int
			fmt.Println("Informe o tempo de prepara e o de entrega")
			fmt.Scan(&preparo, &entrega)
			pedidosPendentes <- Pedido{time.Duration(preparo),
				time.Duration(entrega)}
		case 2:
			var nome string
			fmt.Println("Informe o nome do cozinheiro")
			fmt.Scan(&nome)
			go cozinheiro(nome)

		case 3:
			var nome string
			fmt.Println("Informe o nome do cozinheiro")
			fmt.Scan(&nome)
			go motoboy(nome)
		}
	}
}
