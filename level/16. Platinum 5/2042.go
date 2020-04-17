package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

type fenwick struct {
	number []int
}

func (f *fenwick) add(i, v int) {
	for ; i < len(f.number); i = i | (i + 1) {
		f.number[i] += v
	}
}

func (f *fenwick) get(i int) int {
	s := f.number[i]
	j := i + 1
	j -= j & -j
	for i > j {
		s -= f.number[i-1]
		i -= i & -i
	}
	return s
}

func (f *fenwick) update(i, v int) {
	i -= 1
	v -= f.get(i)
	for c := len(f.number); i < c; i |= i + 1 {
		f.number[i] += v
	}
}

func (f *fenwick) sum(i, j int) (result int) {
	i -= 1
	for j > i {
		result += f.number[j-1]
		j -= j & -j
	}
	for i > j {
		result -= f.number[i-1]
		i -= i & -i
	}
	return
}

func main() {
	defer w.Flush()
	var a, b, c, d, e, f, g int
	fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)
	z := &fenwick{make([]int, a)}

	for i := 0; i < a; i++ {
		fmt.Fscanln(r, &d)
		z.add(i, d)
	}

	for i := 0; i < b+c; i++ {
		fmt.Fscanf(r, "%d %d %d\n", &e, &f, &g)
		if e == 1 {
			z.update(f, g)
		} else if e == 2 {
			fmt.Fprintln(w, z.sum(f, g))
		}
	}
}