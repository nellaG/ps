package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	var num, a, b int
	fmt.Fscanln(reader, &num)
	for i := 0; i < num; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i+1, a, b, a+b)
		writer.Flush()
	}
}
