package main

import "fmt"

const (
	actual = 2024 + iota
	a
	b
	c
	d
)

func main() {

	fmt.Printf("O ano é %v\n", actual)
	fmt.Printf("Proximos anos: \n%v\t%v\t%v\t%v\t", a, b, c, d)

}

/*
	- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
	- Demonstre estes valore
*/
