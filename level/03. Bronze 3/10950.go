package main

import (
	"fmt"
)

func main(){
	var a,b,c int
	fmt.Scanln(&a)

	for i:=0;i<a;i++{
		fmt.Scanln(&b, &c)
		fmt.Println(b+c)
	}
}