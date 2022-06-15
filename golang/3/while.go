package main

import "fmt"

//Não tem While no GO então tem que fazer usando o for
func main() {
	x := 0
	for x < 4 {
		fmt.Print("brabo\n")
		x++
	}
	for {
		if x < 10 {
			x++
			fmt.Println("foda-se PP")
		} else {
			fmt.Println("vá pa puta q pariu")
			break
		}
	}
}
