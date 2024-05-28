package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n, m, t int
	fmt.Fscanln(reader, &t)
	for j := 0; j < t; j++ {
		top := 1
		fmt.Fscanln(reader, &n, &m)
		d := min(n, m-n)
		for i := m; i > m-d; i-- {
			top *= i
		}
		for i := 1; i <= d; i++ {
			top /= i
		}
		fmt.Fprintln(writer, top)
	}
	writer.Flush()
}
