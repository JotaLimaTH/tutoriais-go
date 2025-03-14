package main

import "fmt"

func main(){
	var stringMuitoMassa string = "mario_gomes"
	var indexed = stringMuitoMassa[0]
	fmt.Printf("%s, %v", stringMuitoMassa, indexed)

	for i, v := range stringMuitoMassa {
		fmt.Println(i, v)
	}
	fmt.Printf("\nO tamanho da string 'stringMuitoMassa' é de %d", len(stringMuitoMassa))
}