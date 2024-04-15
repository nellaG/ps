package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nums := make([]int, 5)
	var sum int
	for i := 0; i < 5; i++ {
		fmt.Fscanln(reader, &nums[i])
		sum += nums[i]
	}
	fmt.Println(sum / 5)
	sort.Ints(nums)
	fmt.Println(nums[2])
}
