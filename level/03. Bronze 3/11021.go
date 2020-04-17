package main

import (
	"fmt"
)

func main(){
	var a,b,c int
	fmt.Scanln(&a)

	for i:=1;i<=a;i++{
		fmt.Scanln(&b, &c)
		fmt.Printf("Case #%d: %d\n",i, b+c)
	}
}