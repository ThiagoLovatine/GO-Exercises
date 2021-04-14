package main

import "fmt"

func main() {
	const (
		_ = 2021 + iota
		yearNext1
		yearNext2
		yearNext3
		yearNext4
	)

	fmt.Println(yearNext1, yearNext2, yearNext3, yearNext4)
}
