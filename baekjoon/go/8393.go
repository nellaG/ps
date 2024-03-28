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
	for i := 1; i <= a; i++ {
		b += i
	}
	fmt.Println(b)
}
