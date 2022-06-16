package main

import (
	"fmt"
)

func main() {
	// modulo (%) pega o resto
	// nesse caso se tiver resto ele continua se não printa obs:(3/2 = 2 e sobra 1 já 2/2 = 2 não sobra nada)
	// continue ele segue o fluxo sem fazer nada
	x := 0
	for x <= 20 {
		x++
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
