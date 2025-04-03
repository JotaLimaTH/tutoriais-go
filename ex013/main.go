package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"image/color"
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

	img_2 := image.NewRGBA(image.Rect(0, 0, 25, 25))
	img_2.Set(10, 10, color.NRGBA{R: 255, G: 0, B: 100})

	file_2, _ := os.Create("saida.png")
	defer file_2.Close()

	png.Encode(file_2, img_2)
}