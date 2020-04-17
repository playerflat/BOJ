package main

import "fmt"

func main() {
	var n1 int
	fmt.Scanln(&n1)

	if n1%400 == 0 {
		fmt.Println(1)
	} else if n1%4 == 0 && n1%100 != 0 {
		fmt.Println(1)
	} else{
		fmt.Println(0)
	}
}