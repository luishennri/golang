package main

import "fmt"

func main() {
	s := "hello World"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%s - %t %#U - %#X\n", v, v, v, v)
	}
}
