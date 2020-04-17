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
	var a, b, c, d int
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	fmt.Fscan(r, &c)
	fmt.Fscan(r, &d)

	t := time.Date(2019, time.Month(1), 1, a, b, c, 0, time.UTC)
	t = t.Add(time.Second * time.Duration(d))
	fmt.Fprintf(w, "%d %d %d", t.Hour(), t.Minute(), t.Second())
}