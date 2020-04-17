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
	b := make(map[int]int)
	c := make([]int, 0)

	for i := 1; i <= 30; i++ {
		b[i] = 1
	}

	for i := 0; i < 28; i++ {
		fmt.Fscanln(r, &a)
		delete(b, a)
	}

	for i := range b{
		c = append(c, i)
	}

	if c[0]>c[1]{
		fmt.Fprintf(w, "%d\n%d",c[1],c[0])
	} else{
		fmt.Fprintf(w, "%d\n%d",c[0],c[1])
	}
}