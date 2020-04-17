package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	var a string
	cro := []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}
	fmt.Fscanln(r, &a)

	for _, v := range cro {
		a = strings.Replace(a, v, "0", -1)
	}
	fmt.Fprintln(w, len(a))
}