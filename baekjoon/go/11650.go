package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n, x, y int
	defer writer.Flush()
	fmt.Fscanln(reader, &n)
	cds := make(map[int][]int)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x, &y)
		cds[x+100000] = append(cds[x+100000], y)
	}
	var keys []int
	for k := range cds {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		sort.Ints(cds[k])
		for i := 0; i < len(cds[k]); i++ {
			fmt.Fprintf(writer, "%d %d\n", k-100000, cds[k][i])
		}
	}
}
