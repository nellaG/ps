package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, c, gcd_ int
	var x int
	fmt.Fscanln(reader, &n)
	trees := make([]int, n)
	intervals := make([]int, n-1)
	fmt.Fscanln(reader, &trees[0])
	for i := 1; i < n; i++ {
		fmt.Fscanln(reader, &x)
		trees[i] = x
		intervals[i-1] = x - trees[i-1]
	}
	gcd_ = gcd(intervals[0], intervals[1])
	for i := 2; i < n-1; i++ {
		gcd_ = gcd(gcd_, intervals[i])
	}
	for i := 0; i < n-1; i++ {
		c += intervals[i]/gcd_ - 1
	}
	fmt.Println(c)
}
