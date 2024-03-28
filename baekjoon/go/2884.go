package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscanln(reader, &a, &b)
	switch {
	case b >= 45:
		fmt.Printf("%d %d\n", a, b-45)
	case a == 0:
		fmt.Printf("%d %d\n", 23, b-45+60)
	default:
		fmt.Printf("%d %d\n", a-1, b-45+60)
	}

}
