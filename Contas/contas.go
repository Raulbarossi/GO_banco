package contas

type cc struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *cc) Sacar(valorSaque float64) string {

	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso"

	} else {
		return "Sando insulficiente"
	}
}

func (c *cc) Depositar(valorDep float64) (string, float64) {

	if valorDep > 0 {
		c.Saldo += valorDep
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor do depósito menor que zero", c.Saldo
	}
}

func (c *cc) transferir(valorTransf float64, contaDestino *cc) bool {

	if valorTransf > 0 && valorTransf < c.Saldo {
		c.Saldo -= valorTransf
		contaDestino.Depositar(valorTransf)
		return true
	} else {
		return false
	}

}
