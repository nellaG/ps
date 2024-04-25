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
	var n, m, x int
	fmt.Fscanln(reader, &n)
	cards := make(map[int]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &x)
		cards[x] = 1
	}
	fmt.Fscanln(reader, &m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &x)
		fmt.Fprintf(writer, "%d ", cards[x])
	}
}
