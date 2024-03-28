package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	switch {
	case a > 0 && b > 0:
		fmt.Println(1)
	case a > 0 && b < 0:
		fmt.Println(4)
	case a < 0 && b < 0:
		fmt.Println(3)
	case a < 0 && b > 0:
		fmt.Println(2)
	}
}
