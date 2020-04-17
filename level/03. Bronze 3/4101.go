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
	var a, b int
	for{
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		if a+b == 0{
			break
		}

		if a>b{
			fmt.Fprintln(w, "Yes")
		} else{
			fmt.Fprintln(w, "No")
		}
	}
}