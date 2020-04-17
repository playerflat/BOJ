package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var a int
	fmt.Fscanln(r, &a)
	c := Append(a)

	sort.Ints(c)

	for i := 0; i<a; i++ {
		fmt.Println(c[i])
	}
}

func Append(size int) (s []int){
	var b int
	for i := 0; i < size; i++ {
		fmt.Fscanln(r, &b)
		s = append(s, b)
	}
	return s
}