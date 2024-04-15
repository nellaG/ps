package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanln(reader, &n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })
	fmt.Println(arr[k-1])
}
