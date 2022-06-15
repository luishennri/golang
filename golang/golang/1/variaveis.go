package main

import "fmt"

// := declara variavel fora do bloco, mas dentro do package main no caso
var y = "fala dele"

func main() {

	// := declara variavel dentro do bloco
	x := 10

	fmt.Printf("x: %v, %t \n", x, x)
	fmt.Printf("y: %v, %t\n\n", y, y)

	//	= atribui valor
	x = 45
	fmt.Printf("x: %v, %t \n\n", x, x)

	// := atribui valor se tiver declarando uma varialvel junto
	x, z := 20, 30
	fmt.Printf("x: %v, %t \n", x, x)
	fmt.Printf("z: %v, %t \n", z, z)

}
