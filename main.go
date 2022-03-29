package main

import (
	"fmt"
	"Alura\Banco\contas"
)

func main() {

	cSilvia := contas.cc{"Silvia", 0, 0, 300}
	cGustavo := contas.cc{"Gustavo", 0, 0, 100}

	status := cSilvia.transferir(-200, &cGustavo)
	fmt.Println(status)
	fmt.Println(cSilvia)
	fmt.Println(cGustavo)
}
