package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Compiler) // Compilador
	fmt.Println(runtime.GOOS) // Sistema operacional
	fmt.Println(runtime.GOARCH) // Arquitetura do processador
}
