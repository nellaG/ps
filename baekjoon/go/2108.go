package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, sum, l, c int
	fmt.Fscanln(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a[i])
		sum += a[i]
		l++
	}
	slices.Sort(a)
	fmt.Println(int(math.Round(float64(sum) / float64(l))))
	fmt.Println(a[n/2])

	p := a[0]
	c = 0
	modes := []int{p}
	maxc := 0

	for i := 1; i < n; i++ {
		if p != a[i] {
			p = a[i]
			c = 0
		} else {
			c++
		}
		if c == maxc {
			if len(modes) < 2 {
				modes = append(modes, a[i])
			}
		} else if c > maxc {
			maxc = c
			modes = []int{a[i]}
		}
	}
	mode := modes[0]
	if len(modes) > 1 {
		mode = modes[1]
	}

	fmt.Println(mode)
	fmt.Println(a[n-1] - a[0])
}
