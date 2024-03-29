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
	var from, to, temp int
	result := make([]int, bucket)
	for i := 1; i < bucket+1; i++ {
		result[i-1] = i
	}
	for i := 0; i < line; i++ {
		fmt.Fscanln(reader, &from, &to)
		temp = result[from-1]
		result[from-1] = result[to-1]
		result[to-1] = temp
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
