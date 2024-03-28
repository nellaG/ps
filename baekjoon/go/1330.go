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
	case a == b:
		fmt.Println("==")
	case a < b:
		fmt.Println("<")
	case a > b:
		fmt.Println(">")
	}
}
