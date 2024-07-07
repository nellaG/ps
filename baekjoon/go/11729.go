package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {

	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscanln(reader, &n)
	fmt.Fprintln(writer, int(math.Pow(2, float64(n))-1))
	hanoi(n, 1, 3, 2, writer)
}

func hanoi(n, from, to, tmp int, writer io.Writer) {
	if n == 1 {
		fmt.Fprintf(writer, "%d %d\n", from, to)
		return
	}
	hanoi(n-1, from, tmp, to, writer)
	fmt.Fprintf(writer, "%d %d\n", from, to)
	hanoi(n-1, tmp, to, from, writer)
}
