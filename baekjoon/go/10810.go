package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	var bucket, line int
	fmt.Fscanln(reader, &bucket, &line)
	var from, til, num int
	result := make([]int, bucket)

	for i := 0; i < line; i++ {
		fmt.Fscanln(reader, &from, &til, &num)
		for j := from - 1; j < til; j++ {
			result[j] = num
		}
	}
	for i := 0; i < bucket; i++ {
		fmt.Fprintf(writer, "%d", result[i])
		if i == bucket-1 {
			break
		}
		fmt.Fprintf(writer, " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
