package main

import (
	"fmt"
)

func main() {
	var b, c int

	for {
		fmt.Scanln(&b, &c)
		if b == 0 && c == 0{
			break
		}
		fmt.Printf("%d\n", b+c)
	}
}