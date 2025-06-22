package main

import (
	"fmt"
	"pipetest1/pkg/helpers"
)

// PrintSimple prints, if x*x + x + 1 divided by p in range from 1 to [lim]
func PrintSimple(lim int, p int) {
	for k := 1; k <= lim; k++ {
		target := helpers.X111(k)
		if helpers.DivP(target, p) {
			fmt.Printf("%d is divided by %d\n", target, p)
		}
	}
}

func main() {
	PrintSimple(100, 43)
}
