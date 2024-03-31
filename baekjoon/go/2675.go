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
	var num, repeat int
	var word string
	fmt.Fscanln(reader, &num)
	for i := 0; i < num; i++ {
		fmt.Fscanln(reader, &repeat, &word)
		for _, v := range word {
			fmt.Fprint(writer, strings.Repeat(string(v), repeat))
		}
		fmt.Fprintln(writer)
		writer.Flush()
	}
}
