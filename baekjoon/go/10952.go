package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b int

	for {
		fmt.Fscanln(reader, &a, &b)
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}
}
