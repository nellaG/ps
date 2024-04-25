package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, count int
	var x string
	fmt.Fscanln(reader, &n, &m)
	strs := make(map[string]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x)
		strs[x] = 1
	}
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &x)
		count += strs[x]
	}
	fmt.Println(count)
}
