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
	var n int
	fmt.Fscanln(reader, &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &nums[i])
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, nums[i])
	}
}
