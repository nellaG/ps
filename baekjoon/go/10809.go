package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	var word string
	alb := make([]int, 26)
	fmt.Fscanln(reader, &word)
	for i, v := range word {
		w := v - 97
		if alb[w] == 0 {
			alb[w] = i + 1
		}
	}
	for _, v := range alb {
		fmt.Fprintf(writer, "%d ", v-1)
	}
	writer.Flush()
}
