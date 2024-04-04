package main

import "fmt"

const (
	a = iota
	_ = iota
	b = iota
	c
)

func main() {
	fmt.Println(a, b, c)
}


// Iota cria um numero diferente do anterior
