package main

import "fmt"

// IOTA cria uma sequencia começando de 0
// IOTA segue a sequencia sem precisar digitar o resto do codigo
const (
	a = iota
	b
	c
	//o _ descarta
	_
	e
)

func main() {
	fmt.Println(a, b, c, e)
}
