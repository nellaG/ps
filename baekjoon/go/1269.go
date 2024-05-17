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
	var n, m, c, x int
	fmt.Fscanln(reader, &n, &m)
	map_ := make(map[int]int, n+m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		map_[x] += 1
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &x)
		map_[x] += 1
	}
	for _, v := range map_ {
		if v == 1 {
			c++
		}
	}
	fmt.Fprintln(writer, c)
}
