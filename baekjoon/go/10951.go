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
		_, err := fmt.Fscanln(reader, &a, &b)
		if err != nil {
			break
		}
		fmt.Println(a + b)
	}
}
