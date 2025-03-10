package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println("Valor de x: ", x)
	fmt.Println("Endereço de x: ", &x)
	fmt.Println("Valor de ponteiro p: ", p)
	fmt.Println("Valor apontado pelo ponteiro p: ", *p)

	y := 10
	fmt.Println("Valor de y: ", y)
	dobrarValor(&y)
	fmt.Println("Valor de y depois de ser dobrado: ", y)
}

func dobrarValor(x *int) {
	*x = *x * 2
}