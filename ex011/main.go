package main

import "fmt"

func main() {
	var nome string
	var inteiro int = 5
	var booleano bool = true
	const (
		_ = iota
		a = iota
		_ = iota
		b = iota
		c = iota * 5
	) // Iota
	string_muito_loka := fmt.Sprint("Bom dia. ", "Como vai você?") // Isso vai retornar uma string
	fmt.Print("Isso vai printar sem criar uma nova linha a mais")
	fmt.Println("Isso vai criar uma linha a mais.")
	fmt.Printf("%s Qual o seu nome?\n", string_muito_loka) // Isso vai printar as variáveis que tu formatar
	fmt.Print("Digite seu nome: ")
	fmt.Scanf("%s", &nome) // Isso vai servir como input.
	fmt.Printf("Bom dia, %s! Espero que esteja muito bem. Sua sala é a de número %d, com a placa escrita %t.\n\n", nome, inteiro, booleano)

	fmt.Sscanf("Loucura 12 false", "%s %d %t", &string_muito_loka, &inteiro, &booleano) // Isso vai pegar o conteúdo separado por espaço e vai passar o valor para as variáveis.
	fmt.Printf("%s %d %t", string_muito_loka, inteiro, booleano)

	fmt.Printf("Decimal: %d\nBinário: %b\nHexadecimal: %#x\n", inteiro, inteiro, inteiro)
	fmt.Printf("Constantes iota:\n\n%d\n%d\n%d", a, b, c)
}