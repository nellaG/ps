package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, num int
	fmt.Fscanln(reader, &n)
	counts := make([]int, 10001)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &num)
		counts[num]++
	}
	for i := 1; i < 10001; i++ {
		for j := 0; j < counts[i]; j++ {
			fmt.Fprintln(writer, i)
		}
	}
}
