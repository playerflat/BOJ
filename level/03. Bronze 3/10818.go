package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, b string
	fmt.Scanln(&a)
	in := bufio.NewReader(os.Stdin)
	b, _ = in.ReadString('\n')
	c := strings.Fields(b)
	r1, r2 := minmax(c)
	fmt.Println(r1, r2)

}

func minmax(a []string) (min int, max int) {
	min, _ = strconv.Atoi(a[0])
	max, _ = strconv.Atoi(a[0])
	for _, value := range a {
		if value, _ := strconv.Atoi(value); value < min {
			min = value
		}
		if value, _ := strconv.Atoi(value);value > max {
			max = value
		}
	}
	return min, max
}