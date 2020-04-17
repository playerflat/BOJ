package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a int
	fmt.Fscanln(r, &a)

	for i := 0; i <= a; i++ {
		if a == calc(i) {
			fmt.Println(i)
			break
		}
		if i >= a{
			fmt.Println(0)
		}
	}
}

func calc(a int) (sum int) {
	sum += a
	for {
		sum += a % 10
		a /= 10
		if a==0{
			break
		}
	}
	return
}