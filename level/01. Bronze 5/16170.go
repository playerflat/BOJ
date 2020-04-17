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
	now := time.Now()
	fmt.Fprintln(w, now.Year())
	fmt.Fprintln(w, int(now.Month()))
	fmt.Fprintln(w, now.Day())
}
