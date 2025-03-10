package main

import "fmt"

func main(){
	var intNum int = 120
	fmt.Println(intNum)

	var floatNum float32 = 432.492
	fmt.Println(floatNum)

	fmt.Println(fatorial(5))

	var stringMtoLoka string = "João"
	fmt.Println(stringMtoLoka)
}

func fatorial(numero int) int {
	var resultado int = numero
	for numero > 1 {
		resultado *= numero - 1
		numero --
	}
	return resultado
}