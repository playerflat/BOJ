package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	var n string
	fmt.Fscan(r, &n)

	a := strings.Split(n,"")

	sort.Sort(sort.Reverse(sort.StringSlice(a)))

	b := strings.Join(a,"")
	fmt.Fprintln(w, b)
}