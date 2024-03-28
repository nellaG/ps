package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	var num int
	fmt.Fscanln(reader, &num)
	for i := 1; i <= num; i++ {
		fmt.Fprintf(writer, "%s%s\n", strings.Repeat(" ", num-i), strings.Repeat("*", i))
		writer.Flush()
	}
}
