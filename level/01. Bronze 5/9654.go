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
	fmt.Fprint(w,
`SHIP NAME      CLASS          DEPLOYMENT IN SERVICE
N2 Bomber      Heavy Fighter  Limited    21        
J-Type 327     Light Combat   Unlimited  1         
NX Cruiser     Medium Fighter Limited    18        
N1 Starfighter Medium Fighter Unlimited  25        
Royal Cruiser  Light Combat   Limited    4         
`)
}