package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
var nlist []int

func main() {
	var n, m int
	defer writer.Flush()
	fmt.Fscanln(reader, &n, &m)
	nlist = make([]int, n+1)
	dfs(0, 0, n, m)
	return

}

func dfs(idx int, start int, n int, m int) {
	if idx == m {
		for id, d := range nlist {
			if d == 0 {
				if id != 0 {
					fmt.Fprintf(writer, "\n")
				}
				break
			}
			fmt.Fprintf(writer, "%d ", d)
		}
		return
	}
	for i := start; i < n+1; i++ {
		nlist[idx] = i
		dfs(idx+1, i+1, n, m)
	}
}
