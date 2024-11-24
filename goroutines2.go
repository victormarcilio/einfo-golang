package main

import (
	"fmt"
	"sync"
	"time"
)

type ConexaoDb struct{}

var once sync.Once

func (cdb *ConexaoDb) Conecta() {
	fmt.Println("Conectando ao banco...")
	time.Sleep(3 * time.Second)
	fmt.Println("Conexao realizada...")
}

func (cdb *ConexaoDb) Query() {
	once.Do(cdb.Conecta)
	fmt.Println("Query realizada.")
}

func main() {
	conexao := &ConexaoDb{}
	conexao.Query()
	conexao.Query()
	conexao.Query()
	conexao.Query()
}
