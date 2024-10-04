package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
var nlist []int
var mlist []bool

func main() {
	var n, m int
	defer writer.Flush()
	fmt.Fscanln(reader, &n, &m)
	nlist = make([]int, n+1)
	mlist = make([]bool, n+1)
	dfs(0, n, m)
	return

}

func dfs(idx int, n int, m int) {
	if idx == m {
		for _, d := range nlist {
			if d != 0 {
				fmt.Fprintf(writer, "%d ", d)
			}
		}
		fmt.Fprintf(writer, "\n")
		return
	}
	for i := 1; i < n+1; i++ {
		if !mlist[i] {
			mlist[i] = true
			nlist[idx] = i
			dfs(idx+1, n, m)
			mlist[i] = false
		}
	}
}
