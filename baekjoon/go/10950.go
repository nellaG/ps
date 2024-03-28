package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var num, a, b int
	fmt.Fscanln(reader, &num)

	for i := 0; i < num; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Println(a + b)

	}
}
