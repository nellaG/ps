package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSum(s int) int {
	var sum int
	for s > 0 {
		sum += (s % 10)
		s /= 10
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	for i := 1; i <= n; i++ {
		if getSum(i)+i == n {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(0)
}
