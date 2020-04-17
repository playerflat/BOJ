package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a, b int
	var c string
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscanf(r, "%d %s\n", &b, &c)
		oct, _ := strconv.ParseInt(c, 8, 32)
		dec, _ := strconv.ParseInt(c, 10, 32)
		hex, _ := strconv.ParseInt(c, 16, 32)
		fmt.Fprintf(w, "%d %d %d %d", b, oct, dec, hex)
		fmt.Fprintln(w, "")
	}
}