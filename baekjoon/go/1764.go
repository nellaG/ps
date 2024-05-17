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
	defer writer.Flush()
	var n, m, j int
	var x string
	fmt.Fscanln(reader, &n, &m)
	map_ := make(map[string]int, n+m)
	ans := make([]string, 0)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x)
		map_[x] += 1
	}
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &x)
		map_[x] += 1
	}
	for k, v := range map_ {
		if v == 2 {
			ans = append(ans, k)
			j++
		}
	}
	fmt.Fprintln(writer, len(ans))
	sort.Strings(ans)
	for _, x := range ans {
		fmt.Fprintln(writer, x)
	}

}
