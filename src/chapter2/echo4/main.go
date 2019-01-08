package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "skip newline symbol")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
