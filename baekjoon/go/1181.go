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
	var n int
	var before string
	defer writer.Flush()
	fmt.Fscanln(reader, &n)
	strs := make([]string, n)
	for i := range strs {
		fmt.Fscanln(reader, &strs[i])
	}

	sort.Slice(strs, func(i, j int) bool {
		if len(strs[i]) == len(strs[j]) {
			return strs[i] < strs[j]
		}
		return len(strs[i]) < len(strs[j])
	})
	for _, k := range strs {
		if before != k {
			fmt.Fprintf(writer, "%s\n", k)
		}
		before = k
	}
}
