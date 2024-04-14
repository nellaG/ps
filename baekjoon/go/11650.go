package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type vertex struct{ x, y int }

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	defer writer.Flush()
	fmt.Fscanln(reader, &n)
	cds := make([]vertex, n)
	for i := range cds {
		fmt.Fscanln(reader, &cds[i].x, &cds[i].y)
	}

	sort.Slice(cds, func(i, j int) bool {
		if cds[i].x == cds[j].x {
			return cds[i].y < cds[j].y
		}
		return cds[i].x < cds[j].x
	})
	for _, k := range cds {
		fmt.Fprintf(writer, "%d %d\n", k.x, k.y)
	}
}
