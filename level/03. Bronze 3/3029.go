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
	var h1, h2, m1, m2, s1, s2 int
	var wait, throw time.Time
	var flag bool
	fmt.Fscanf(r, "%d:%d:%d\n", &h1, &m1, &s1)
	fmt.Fscanf(r, "%d:%d:%d\n", &h2, &m2, &s2)

	if h1==h2 && m1==m2 && s1==s2 {
		fmt.Fprintln(w, "24:00:00")
	} else {
		if h1 > h2 || (h1 == h2 && m1 > m2) || (h1 == h2 && m1 == m2 && s1 > s2) {
			flag = true
		}

		if flag {
			wait = time.Date(2019, time.Month(1), 1, h1, m1, s1, 0, time.Local)
			throw = time.Date(2019, time.Month(1), 2, h2, m2, s2, 0, time.Local)
			a := wait.Sub(throw) / time.Second
			fmt.Printf("%02d:%02d:%02d", Abs(a/3600), Abs(a%3600/60), Abs(a%3600%60))
		} else {
			wait = time.Date(2019, time.Month(1), 1, h1, m1, s1, 0, time.Local)
			throw = time.Date(2019, time.Month(1), 1, h2, m2, s2, 0, time.Local)
			a := throw.Sub(wait) / time.Second
			fmt.Printf("%02d:%02d:%02d", a/3600, a%3600/60, a%3600%60)
		}
	}
}

func Abs(a time.Duration) time.Duration {
	if a < 0 {
		a *= -1
	}
	return a
}
