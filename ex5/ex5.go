package main

import "fmt"

type fuinha int

var x fuinha

var y int

func main() {

	fmt.Println("Valor Inicial de X: ", x)
	fmt.Printf("Tipo da variavel X: %T\n", x)
	x = 42
	fmt.Println("Valor de X: ", x)

	y = int(x)
	fmt.Println("Valor de Y: ", y)
	fmt.Printf("Tipo da variavel Y: %T\n", y)
}

/*
Utilizando a solução do exercicio anterior:
	1. Em package-level scope, utilizando a palavra var, crie uma variavel com o identificar "y". O tipo dessa variavel deve ser o tipo subjacente do tipo que você criou no exercicio anterior.
	2.Na função main
		1. Isto já deve estar feito:
			1. Demontre o valor da variavel "x"
			2. Demonstre o tipo da variavel "x"
			3. Atribua o valor 42 à variavel "x" utilizando o operador "="
			4.Demonstre o valor da variavel "x"
		2. Agora Faça também
			1.Utilize conversão para transformar o tipo do valor da variavel "X" em seu tipo subjascente e, utilizando o operador "=", atribua o valor de "x" a "y"
			2. Demonstre o valor de "y"
			3. Demonstre o typo de "y"
*/
