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

func (f *fenwick) get(i int) int{
	s := f.number[i]
	j := i+1
	j -= j & -j
	for i > j {
		s -= f.number[i-1]
		i -= i & -i
	}
	return s
}

func (f *fenwick) update(i, v int) {
	i-=1
	v -= f.get(i)
	for c := len(f.number); i < c; i |= i + 1 {
		f.number[i] += v
	}
}

func (f *fenwick) sum(i, j int) (result int) {
	i-=1
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
	fmt.Fscanf(r, "%d %d\n", &a, &b)
	z := &fenwick{make([]int, a)}
	for i := 0; i <a; i++ {
		fmt.Fscanf(r, "%d ", &c)
		z.add(i, c)
	}
	fmt.Fscanln(r, "")

	for i := 0; i <b; i++ {
		fmt.Fscanf(r, "%d %d %d %d\n", &d, &e, &f, &g)
		if d < e {
			fmt.Fprintln(w, z.sum(d, e))
		} else {
			fmt.Fprintln(w, z.sum(e, d))
		}
		z.update(f, g)
	}
}