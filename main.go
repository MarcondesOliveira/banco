package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contadaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(-200, &contadaSilvia)

	fmt.Println(status)

	fmt.Println(contadaSilvia)
	fmt.Println(contaDoGustavo)
}
