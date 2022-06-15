package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {

	fmt.Println("Binary\t\t\t\t\t\t\t\t\t\tDecimal")
	fmt.Printf("%b\t\t\t\t\t\t\t\t\t%d\n", KB, KB)
	fmt.Printf("%b\t\t\t\t\t\t%d\n", MB, MB)
	fmt.Printf("%b\t\t\t\t%d\n", GB, GB)
	fmt.Printf("%b\t%d", TB, TB)

}
