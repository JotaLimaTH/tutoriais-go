package main

import (
	"fmt"
	"runtime"
)

func main() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	// Constantes do package run time
	fmt.Println(runtime.Compiler) // Compilador
	fmt.Println(runtime.GOOS) // Sistema operacional
	fmt.Println(runtime.GOARCH) // Arquitetura do processador
	fmt.Println("Número de CPUs: ", runtime.NumCPU())
	fmt.Println("Memória alocada: ", memStats.Alloc, "bytes") // Retorna em bytes
}
