package main

import "fmt"

func NewConta(titular string, numero string, saldo float64) *conta {
	return &conta{
		Titular: titular,
		Numero:  numero,
		saldo:   saldo,
	}
}

type conta struct {
	Titular string
	Numero  string
	saldo   float64
}

func novaConta(titular string) conta {
	return conta{
		Titular: titular,
		Numero:  "52104",
		saldo:   0.0,
	}
}

func (c *conta) TitularDaConta() string {
	return c.Titular
}

func (c *conta) AlterarTitular(titular string) {
	c.Titular = titular
}

func main() {
	conta_1 := conta{Titular: "Super Homem", Numero: "123456", saldo: 0.0}

	fmt.Printf("%v\n", conta_1.Titular)

	t := NewConta("joao", "123", 457)
	fmt.Println(t.TitularDaConta())
	t.AlterarTitular("jorge")
	fmt.Println(t.Titular)
}
