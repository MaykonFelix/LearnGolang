package main

import (
	"fmt"
)

type fuinha int

var x fuinha

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

/*
- Crie um tipo . O Tipo subjacente deve ser int.
- Crie uma variável para este tipo, como o identificador "x", utilizando a palavra-chave var.
- Na função main:

	1. Demonstre o valor da variável "x"
	2. Demonstre o tipo da variável "x"
	3. Atribua 42 à variável "x" utilizando o operador "="
	4. Demonstre o valor da variável "x"

*/
