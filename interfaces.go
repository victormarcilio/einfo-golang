package main

import (
	"fmt"
	"time"
)

type Desenhavel interface {
	Desenhar()
	Colorir()
	GetNome() string
}

type Triangulo struct{}
type Gato struct{}

func (t Triangulo) Desenhar() {
	fmt.Println("Iniciando desenho...")
	fmt.Println("Desenhando primeiro lado...")
	fmt.Println("Desenhando segundo lado...")
	fmt.Println("Desenhando terceiro lado...")
}

func (t Triangulo) Colorir() {
	fmt.Println("Mexendo a tinta...")
	fmt.Println("Iniciando a pintura...")
	time.Sleep(2 * time.Second)
	fmt.Println("Pintura finalizada")
}

func (_ Gato) Desenhar()       {}
func (_ Gato) Colorir()        {}
func (_ Gato) GetNome() string { return "GATO" }

func Desenha(d Desenhavel) {
	fmt.Println("A desenhar", d.GetNome(), "...")
	d.Desenhar()
	fmt.Println("A colorir", d.GetNome(), "...")
	d.Colorir()
}

func (t Triangulo) GetNome() string {
	return "TRIANGULO"
}

func (t Triangulo) Rotacionar() {
	fmt.Println("Rotacionando...")
}

func f(x int) Desenhavel {
	var t *Triangulo
	fmt.Println("t é nulo?", t == nil)
	if x == 5 {
		fmt.Println("estou retornando t")
		return Triangulo{}
	}
	if x == 7 {
		return Gato{}
	}
	fmt.Println("Nao entrei no if")
	return nil
}

func main() {
	t := Triangulo{}
	Desenha(t)

	fmt.Println("Valor retornado", f(5))
	fmt.Println("Valor retornado", f(6))
	fmt.Println("Valor retornado comparado com nulo", f(5) == nil)
	fmt.Println("Valor retornado comparado com nulo", f(6) == nil)

	ret := f(5)
	cast := ret.(Triangulo) // funciona porque quando o parametro é 5, o tipo retornado é um triangulo

	ret = f(6)
	cast, ok := ret.(Triangulo) //type assertion, ok será verdadeiro se o valor contido em ret for mesmo triangulo
	if ok {
		cast.Rotacionar()
	} else {
		fmt.Print("o cast falhou")
	}
	v := 9
	ret = f(v)
	switch ret := ret.(type) { // cria um novo ret que será do tipo do case que entrar
	case Gato: // aqui ret é um gato
		fmt.Println("Meow", ret.GetNome())
	case Triangulo: // aqui ret é um triangulo
		fmt.Println("Triangulo", ret.GetNome())
	default:
		fmt.Println("nil")
	}
}
