package main

import "fmt"

func main() {

	x := "ZIKA"
	switch {
	case x == "brabo":
		fmt.Println("EU SOU BRRAABO!!")
	case x == "ZIKA":
		fmt.Println("EU SOU ZIKA!!")
	default:
		fmt.Println("EU NÃO SOU BRABO :(")
	}
}
