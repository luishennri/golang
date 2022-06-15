package main

import "fmt"

func main() {

	//é um loop dentro de outro loop, cada ciclo do loop menor vai ocorrer para cada numero diferente do loop maior
	for horas := 0; horas <= 24; horas++ {
		fmt.Println("hora:", horas)
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Print(" ", minutos)
		}
		fmt.Print("\n")
	}
}
