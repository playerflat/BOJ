package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var n, k, v, count int
	coin := make([]int, 0)
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &k)
	for i:=0;i<=n;i++{
		fmt.Fscanln(r, &v)
		coin = append(coin, v)
	}

	for{
		if k==0{
			break
		}
		count+=k/coin[n]
		k%=coin[n]
		n--
	}
	fmt.Fprintln(w,count)
}