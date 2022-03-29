package contas

import "Alura/Banco/clientes"

type CP struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *CP) Sacar(valorSaque float64) string {

	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"

	} else {
		return "Sando insulficiente"
	}
}
func (c *CP) Depositar(valorDep float64) (string, float64) {

	if valorDep > 0 {
		c.saldo += valorDep
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}
func (c *CP) ObterSaldo() float64 {
	return c.saldo
}
