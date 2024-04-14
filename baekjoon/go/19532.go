package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c, d, e, f int
	fmt.Fscan(reader, &a, &b, &c, &d, &e, &f)
	for x := -999; x <= 999; x++ {
		for y := -999; y <= 999; y++ {
			if a*x+b*y == c && d*x+e*y == f {
				fmt.Printf("%d %d\n", x, y)
			}
		}
	}
}
