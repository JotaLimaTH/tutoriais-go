package main

import "fmt"

type Pessoa struct{
	Nome string
	Idade int
	id int
}

func newPessoa(nome string, idade int, idPessoa int) Pessoa{
	pessoa := Pessoa{Nome: nome, Idade: idade, id: idPessoa}
	return pessoa
}

func main(){
	var joao Pessoa // Inicialização sem valor
	fmt.Println(joao.Nome, joao.Idade, joao.id) // Sem valor ainda

	joao.Nome = "João"
	joao.Idade = 20
	joao.id = 123345
	// Atribuição de valor um por um
	fmt.Println(joao.Nome, joao.Idade, joao.id)

	var pedro Pessoa = Pessoa{Nome: "Pedro", Idade: 18, id: 34567,} // Inicialização com atribuição
	fmt.Println(pedro.Nome, pedro.Idade, pedro.id)

	var maria Pessoa = newPessoa("Maria", 22, 123445) // Inicialização com função construtora
	fmt.Println(maria.Nome, maria.Idade, maria.id)
}