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
	for {
		var a, sum int
		fmt.Fscanln(r, &a)

		if a == -1 {
			break
		}

		for i := 1; i < a; i++ {
			if a%i == 0 {
				sum += i
			}
		}
		if sum == a {
			fmt.Fprintf(w, "%d = ", a)
			for i := 1; i <= a/2; i++ {
				if a%i == 0 {
					fmt.Fprintf(w, "%d", i)
					if i != a/2 {
						fmt.Fprintf(w, " + ")
					}
				}
			}
			fmt.Fprintln(w, "")
		} else {
			fmt.Fprintf(w, "%d is NOT perfect.\n", a)
		}
	}
}