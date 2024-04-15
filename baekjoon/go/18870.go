package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	var cdss []int
	before := -9999999999
	fmt.Fscanln(reader, &n)
	cds := make([]int, n)
	cdscp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &cds[i])
		cdscp[i] = cds[i]
	}
	sort.Ints(cds)
	for _, k := range cds {
		if before != k {
			cdss = append(cdss, k)
		}
		before = k
	}
	for _, k := range cdscp {
		idx, _ := slices.BinarySearch(cdss, k)
		fmt.Fprintf(writer, "%d ", idx)
	}

}
