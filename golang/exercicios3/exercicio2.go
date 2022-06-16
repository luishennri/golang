package main

import "fmt"

func main() {

	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for l := 1; l <= 3; l++ {
			fmt.Printf("%#U\n", x)
		}
	}
}
