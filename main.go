package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(5000)

	fmt.Println(contaDoDenis.ObterSaldo())
}
