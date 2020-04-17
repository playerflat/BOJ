package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a1, a2, a3, b1, b2, b3 int
	fmt.Fscanf(r, "%d:%d:%d\n", &a1, &a2, &a3)
	fmt.Fscanf(r, "%d:%d:%d\n", &b1, &b2, &b3)

	atime := time.Date(2019, time.Month(12), 17, a1, a2, a3, 0, time.Local)
	var btime time.Time
	if b1 > a1 || (b1 == a1 && b2 > a2) || (b1 == a1 && b2 == a2 && b3 > a3) {
		btime = time.Date(2019, time.Month(12), 17, b1, b2, b3, 0, time.Local)
	} else {
		btime = time.Date(2019, time.Month(12), 18, b1, b2, b3, 0, time.Local)
	}
	a := btime.Sub(atime)
	fmt.Fprintf(w, "%02.f:%02.f:%02.f", math.Floor(a.Hours()), math.Floor(math.Mod(a.Minutes(),60)), math.Floor(math.Mod(a.Seconds(),60)))
}