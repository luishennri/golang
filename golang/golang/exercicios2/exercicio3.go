package main

import "fmt"

func main() {

	x := 10

	fmt.Printf("%d %b %x\n\n", x, x, x)

	y := x << 1

	fmt.Printf("%d %b %x", y, y, y)
}
