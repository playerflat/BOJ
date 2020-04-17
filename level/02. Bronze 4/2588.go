package main

import "fmt"

func main(){
	var n1, n2 int
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Println(n2%10*n1)
	fmt.Println(n2/10%10*n1)
	fmt.Println(n2/10/10*n1)
	fmt.Println(n2*n1)
}