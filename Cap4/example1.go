package main

import "fmt"

func main() {
	a := "a"
	b := "é"
	c := "u333"

	fmt.Printf("%v %v %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v %v %v", d, e, f)

}


// Exemplo de Codificação UTF-8, para mostrar que são tamanhos diferentes alocados na memoria