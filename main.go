package main

import "fmt"

type cc struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func main() {
	cGuilherme := cc{"Guilherme", 3125, 22335513, 1055.43}
	cBruna := cc{"Bruna", 2531, 33551236, 135.57}

	fmt.Println(cGuilherme)
	fmt.Println(cBruna)

}
