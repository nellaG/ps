package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	var name, p string
	fmt.Fscanln(reader, &n, &m)
	maps := make(map[string]int, n)
	slices := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &name)
		maps[name] = i + 1
		slices[i] = name
	}
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &p)
		x, err := strconv.Atoi(p)
		if err == nil {
			fmt.Fprintln(writer, slices[x-1])
		} else {
			fmt.Fprintln(writer, maps[p])
		}
	}
}
