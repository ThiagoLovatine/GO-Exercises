package main

import "fmt"

func main() {
	var a int = 20

	fmt.Printf("%d %b %#x\n", a, a, a)

	var b = a << 1

	fmt.Printf("%d %b %#x", b, b, b)
}
