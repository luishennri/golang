package main

import "fmt"

const x = 10

//var y = 10

//var c int
var d float64

func main() {
	//nao funciona se a variavel float receber o valor de uma variavel int e vice e versa
	//d = y
	d = x
	fmt.Println(d)
	//já a const nao tem o type definido entao ela consegue usar o valor atribuido
}
