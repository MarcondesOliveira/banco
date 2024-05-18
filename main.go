package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDeExemplo := contas.ContaCorrente{}
	contaDeExemplo.Depositar(100)

	fmt.Println(contaDeExemplo.ObterSaldo())
}
