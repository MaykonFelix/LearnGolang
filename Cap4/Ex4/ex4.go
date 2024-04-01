package main

import "fmt"

func main() {
	var x int = 200

	fmt.Printf("Em decimal: %d\n", x)
	fmt.Printf("Em Binario: %b\n", x)
	fmt.Printf("Em Hexadecimal: %#x\n", x)

	y := x << 1

	fmt.Printf("Em decimal: %d\n", y)
	fmt.Printf("Em Binario: %b\n", y)
	fmt.Printf("Em Hexadecimal: %#x\n", y)

}

/*
Crie um programa que :
	- Atribua um valor int a uma variável
	- Demonstre este valor em decima, binario e hexadecimal
	- Desloque os bits dessa variável 1 para a esquerda e atribua o resultado a outra variável
	- Demonstre esta outra variável em decimal, binário e hexadecimal
*/
