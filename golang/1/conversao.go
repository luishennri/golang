package main

import "fmt"

// criando o type
type hotdog int

var b hotdog = 1011

func main() {
	// variaveis de tipos diferentes
	x := 10
	fmt.Printf("x: %v, %t\n", x, x)
	fmt.Printf("b: %v, %t\n", b, b)

	//para atribuir valor para variaveis diferente é necessario converter o tipo
	x = int(b)
	fmt.Printf("x: %v, %t\n", x, x)
}
