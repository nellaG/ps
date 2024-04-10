package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPrime(i int) bool {
	if i == 1 {
		return false
	}
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var m, n, min_, sum int
	first := true
	fmt.Fscanln(reader, &m)
	fmt.Fscanln(reader, &n)
	for i := m; i <= n; i++ {
		if isPrime(i) {
			if first {
				min_ = i
				first = false
			}
			sum += i
		}
	}
	if sum == 0 {
		fmt.Println(-1)
	} else {
		fmt.Printf("%d\n%d", sum, min_)
	}
}
