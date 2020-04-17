package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a, b, c int
	fmt.Fscan(r, &a, &b, &c)

	t := time.Date(2019, 1, 1, a, b, 0, 0, time.UTC)
	t = t.Add(time.Minute * time.Duration(c))

	fmt.Fprintln(w, t.Hour(), t.Minute())
}