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
	fmt.Fprintln(w, `     /~\
    ( oo|
    _\=/_
   /  _  \
  //|/.\|\\ 
 ||  \ /  ||
============
|          |
|          |
|          |`)
}