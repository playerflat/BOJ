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
	var a, b int
	fmt.Fscanf(r, "%d %d\n", &a, &b)

	t := time.Date(2009, time.Month(b), a,12,00,00,00,time.FixedZone("KST",9*60*60))
	fmt.Fprintln(w, t.Weekday())
}