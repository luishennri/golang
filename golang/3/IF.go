package main

import "fmt"

func main() {

	// o ! inverte, é um não, nega a condição
	x := 200

	if x < 100 {
		fmt.Println("X é menor que 100")
	} else if !(x > 100) {
		fmt.Println("X é maior que 100")
	} else {
		fmt.Println("X é doido")
	}
}
