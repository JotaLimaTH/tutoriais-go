package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func main(){
	file, err := os.Open("images/poltrona.png")
	if err != nil {
		fmt.Println("Erro ao abrir a imagem: ", err)
		return
	}
	defer file.Close()
	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Println("Erro ao decodificar a imagem: ", err)
		return
	}
	fmt.Println("Formato da imagem: ", format)
	fmt.Printf("Dimensões: %dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())
}