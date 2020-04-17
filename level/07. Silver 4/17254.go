package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

type keyboard struct {
	num  int
	time int
	word string
}

type keyboards []keyboard

func (k keyboards) Len() int {
	return len(k)
}

func (k keyboards) Less(i, j int) bool {
	if k[i].time != k[j].time {
		return k[i].time < k[j].time
	} else {
		return k[i].num < k[j].num
	}
}

func (k keyboards) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func main() {
	defer w.Flush()
	var a, b, c, d int
	var e string
	f := make(keyboards, 0)
	fmt.Fscanf(r, "%d %d\n", &a, &b)
	for i := 0; i < b; i++ {
		fmt.Fscanf(r, "%d %d %s\n", &c, &d, &e)
		f = append(f, keyboard{num: c, time: d, word: e})
	}

	sort.Sort(f)
	for i, _ := range f {
		fmt.Fprint(w, f[i].word)
	}
}