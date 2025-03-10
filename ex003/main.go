package main

import "fmt"

func main() {
	const constanteMtoLoka = "Sabe mto"
	const pi float32 = 3.14
	fmt.Println(constanteMtoLoka, "\n", pi)

	var resultado, resto int = divisaoInteira(50, 2)

	fmt.Printf("O resultado dessa divisão é %v, e o resto é %v.\n", resultado, resto) //Caso se queira printar coisas com variáveis, melhor usar o Printf. Aqui é parecido com C, sendo que o %v é o curinga, só que, obviamente, é pouco performático.
}

func divisaoInteira (numerador int, denominador int) (int, int) {
	var resultado int = numerador / denominador
	var resto int = numerador % denominador
	return resultado, resto
}