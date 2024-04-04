package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, comb int
	fmt.Fscanln(reader, &n)
	comb = 1
	for i := 0; i < n; i++ {
		comb += 6 * i
		if n <= comb {
			fmt.Println(i + 1)
			break
		}

	}
}
