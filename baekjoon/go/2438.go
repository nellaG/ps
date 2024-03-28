package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	var num int
	fmt.Fscanln(reader, &num)
	for i := 1; i <= num; i++ {
		for j := 0; j < i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "\n")
		writer.Flush()
	}
}
