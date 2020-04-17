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
	var a, b, sum int
	c := make([]int, 0, 1)
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscanln(r, &b)
		if b ==0{
			c = c[:len(c)-1]
		} else{
			c = append(c, b)
		}
	}
	for _, v := range c{
		sum+=v
	}
	fmt.Fprintln(w, sum)
}