package main

import "fmt"

type loko int

var x loko

func main() {

	fmt.Println(x)
	fmt.Printf("%t\n", x)
	x = 42
	fmt.Println(x)

}
