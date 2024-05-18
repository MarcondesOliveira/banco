package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Salvo insuficiente"
	}
}

func main() {
	contadaSilvia := ContaCorrente{}
	contadaSilvia.titular = "Silvia"
	contadaSilvia.saldo = 500

	fmt.Println(contadaSilvia.saldo)

	fmt.Println(contadaSilvia.sacar(400))

	fmt.Println(contadaSilvia.saldo)

}
