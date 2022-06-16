package main

import "fmt"

func main() {
	x := 4
	// switch é igual o IF só que diferente
	switch {
	case x < 5:
		fmt.Println("x é menor que 5")
	case x == 5:
		fmt.Println("x é  5")
	case x > 5:
		fmt.Println("x é maior que 5")
	}

	melhoresduo := "joao"

	switch melhoresduo {
	case "rogerio", "dudz":
		fmt.Println("joao ")
	case "luis", "brabo":
		fmt.Println("luis é o brabo")
	case "pedro", "joao":
		fmt.Println("pedro e joao")
		// fallthrough é se esse case for verdade o proximo também vai ser
		fallthrough
	// default é o else do IF
	default:
		fmt.Println("mentira, só tem bot")
	}
}
