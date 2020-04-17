package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a, b, sum int
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscanln(r, &b)
		if b == 100 {
			sum += b
		} else if c := b % 10; c == 0 || c == 6 || c == 9 {
			sum += 9
			if e := b - c; e == 60 || e == 90 {
				sum += 90
			} else {
				sum += b - c
			}
		} else if c := b / 10 * 10; c == 60 || c == 90 {
			sum += 90
			sum += b % 10
		} else {
			sum += b
		}
	}
	fmt.Fprintln(w, math.Round(float64(sum)/float64(a)))
}