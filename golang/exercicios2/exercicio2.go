package main

import "fmt"

func main() {

	a := (10 == 10)
	b := (10 != 9)
	c := (10 < 20)
	d := (10 > 5)
	e := (10 <= 7)
	f := (10 >= 12)

	fmt.Println(a, b, c, d, e, f)

}
