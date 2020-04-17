package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a, sum int
	var flag bool
	as := make([]int,0)

	for i:=0;i<9;i++{
		fmt.Fscan(r, &a)
		as = append(as, a)
		sum += a
	}

	for i:=0;i<len(as);i++{
		sum-=as[i]
		for j:=i+1;j<len(as);j++{
			sum-=as[j]
			if sum==100{
				flag=true
				as[j]=0
				as[i]=0
				break
			} else{
				sum+=as[j]
			}
		}
		if flag{
			break
		} else{
			sum+=as[i]
		}
	}

	sort.Ints(as)

	for i:=2;i<9;i++ {
		fmt.Fprintln(w, as[i])
	}
}