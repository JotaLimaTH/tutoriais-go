package main

import "fmt"

func main() {
	string_muito_loka := fmt.Sprint("Bom dia. ", "Como vai você?") // Isso vai retornar uma string
	fmt.Print("Isso vai printar sem criar uma nova linha a mais")
	fmt.Println("Isso vai criar uma linha a mais.")
	fmt.Printf("%s Qual o seu nome?\n", string_muito_loka) // Isso vai printar as variáveis que tu formatar
}