package contas

import "Alura/Banco/clientes"

type CC struct {
	Titular        clientes.Titular
	Agencia, Conta int
	saldo          float64
}

func (c *CC) Sacar(valorSaque float64) string {

	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"

	} else {
		return "Sando insulficiente"
	}
}

func (c *CC) Depositar(valorDep float64) (string, float64) {

	if valorDep > 0 {
		c.saldo += valorDep
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

func (c *CC) Transferir(valorTransf float64, contaDestino *CC) bool {

	if valorTransf > 0 && valorTransf < c.saldo {
		c.saldo -= valorTransf
		contaDestino.Depositar(valorTransf)
		return true
	} else {
		return false
	}

}

func (c *CC) ObterSaldo() float64 {
	return c.saldo
}
