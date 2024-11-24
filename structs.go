package main

import "fmt"

type Motor struct {
	ligado bool
}

func (m *Motor) Ligar() {
	m.ligado = true
	fmt.Println("ligando motor...")
}

func (m *Motor) Desligar() {
	m.ligado = false
	fmt.Println("desligando motor...")
}

type Carro struct {
	Motor
	placa string
	marca string
}

type Aluno struct {
	Nome  string
	notas [3]int
}

type Ponto struct {
	X int
	Y int
}

func (p *Ponto) dx(dist int) {
	p.X += dist
}

func (p *Ponto) dy(dist int) {
	p.Y += dist
}

type Pixel struct {
	Cor   string
	Ponto // Dessa forma eu consigo me referir a X e Y como se tivessem sido declarados nessa estrutura
}

func (p *Pixel) Pintar(novaCor string) {
	if p == nil {
		fmt.Println("erro: chamou pintar num objeto nulo")
		return
	}
	p.Cor = novaCor
}

type Pixel2 struct {
	Cor string
	p   Ponto // Dessa forma eu preciso acessar primeiro o campo p para poder ter acesso a X e Y
}

func main() {
	var a Aluno
	fmt.Println(a) // todos os campos com valor 'zero'
	a.notas = [3]int{7, 8, 5}
	a.Nome = "Joao"

	b := Aluno{"Joao", [3]int{7, 8, 5}} // inicialização com literal deve conter todos os campos na mesma ordem da declaração da estrutura

	c := Aluno{
		Nome: "Beatriz", // Ao nomear campo a campo, eu posso inicializar parcialmente a estrutura, os demais campos permanecerão com o valor 'zero'
	}

	fmt.Println(a == b)
	fmt.Println(c)

	pix1 := Pixel{"Amarelo", Ponto{25, 50}}
	pix2 := &Pixel2{"Azul", Ponto{50, 50}}

	fmt.Println(pix1, pix2)
	pix1.Pintar("Branco")
	var p *Pixel
	p.Pintar("Preto")
	fmt.Println(pix1, pix2)
	fmt.Println(pix1.X)
	fmt.Println(pix2.p.X)

	ponto := Ponto{5, 10}
	fmt.Println(ponto)
	ponto.dx(30)
	ponto.dy(-20)
	fmt.Println(ponto)

	pix1.dx(5)
	pix1.dy(15)

	pix2.p.dx(100)
	fmt.Println(pix1, pix2)
}
