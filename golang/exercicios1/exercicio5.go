package main

import "fmt"

type loko int

var x loko
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%t\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%t\n", y)
}
