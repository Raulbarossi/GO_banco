package main

import (
	"Alura/Banco/contas"
	"fmt"
)

func main() {

	fmt.Println("")
	contaBruno := contas.CP{}
	contaBruno.Depositar(150)
	fmt.Println(contaBruno.ObterSaldo())
	PagarBoleto(&contaBruno, 75)
	fmt.Println(contaBruno.ObterSaldo())
	fmt.Println("================")
	contaBruna := contas.CC{}
	contaBruna.Depositar(500)
	fmt.Println(contaBruna.ObterSaldo())
	PagarBoleto(&contaBruna, 450)
	fmt.Println(contaBruna.ObterSaldo())
	fmt.Println("")
}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
