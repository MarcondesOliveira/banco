package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoTroll := ContaCorrente{
		titular:       "Troll",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.05,
	}

	contaDoZezinho := ContaCorrente{
		"Zezinho",
		222,
		111222,
		200,
	}

	fmt.Println(contaDoTroll)
	fmt.Println(contaDoZezinho)

	var contaDoColosso *ContaCorrente
	contaDoColosso = new(ContaCorrente)
	contaDoColosso.titular = "Colosso"
	contaDoColosso.saldo = 500

	fmt.Println(*contaDoColosso)

}
