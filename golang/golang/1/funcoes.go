package main

import "fmt"

// := declara variavel fora do bloco, mas dentro do package main no caso
var z = 55

func main() {
	// := declara variavel dentro do bloco
	y := 20
	// passando o valor para a outra função
	printarN(y)
}
func printarN(x int) {
	fmt.Println(x)
	fmt.Println(z)
}
