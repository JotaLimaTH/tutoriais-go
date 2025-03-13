package main

import "fmt"

func main(){
	var intArr [3]int32 //Inicialização vazia
	intArr[1] = 12
	intArr[2] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0], &intArr[1], &intArr[2])

	var intArr2 [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArr2) //Inicialização com valores

	intArr3 := [3]int32{200, 400, 600}
	fmt.Println(intArr3) //Inicialização direta 

	var intSlice []int32 = []int32{4,5,6}
	intSlice = append(intSlice, 8)
	fmt.Println(intSlice)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"João":23, "Bento":19}
	fmt.Println(myMap2)
}