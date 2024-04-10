package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	var x = make([]int, n)
	var y = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}
	sort.Ints(x)
	sort.Ints(y)
	fmt.Println((x[n-1] - x[0]) * (y[n-1] - y[0]))
}
