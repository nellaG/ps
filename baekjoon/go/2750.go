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
	fmt.Fscanln(reader, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &arr[i])
	}
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
}
