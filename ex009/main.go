package main

import "fmt"

func main(){
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("O valor de p é %v", p)
	fmt.Printf("\nO endereço de p é %v", &p) // O endereço de p é diferente do valor de p
	fmt.Printf("\nO valor para a qual p aponta é: %v", *p)
	fmt.Printf("\nO valor de i é: %v\n", i)
	*p = 5
	fmt.Println("O valor de *p foi alterado.")
	fmt.Printf("O valor de p é %v", p)
	fmt.Printf("\nO endereço de p é %v", &p)
	// Tanto o valor de p quanto seu endereço permanecem o mesmo, mas o valor para o qual apontam mudou
	fmt.Printf("\nO valor para a qual p aponta é: %v", *p)
	// seloko bicho
}