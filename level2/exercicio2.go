package main

import "fmt"

func main() {
	a := (true != false)
	b := (1 == 1)
	c := (10 > 1)
	d := (10 < 1)
	e := (10 >= 10)
	f := (10 <= 9)
	fmt.Println(a, b, c, d, e, f)
}
