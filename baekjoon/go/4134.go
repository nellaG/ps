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
	for j := 2; j*j <= i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, x, gg int
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x)
		if x < 3 {
			fmt.Fprintln(writer, 2)
			continue
		}
		gg = x + 1
		if isPrime(x) {
			fmt.Fprintln(writer, x)
			continue
		}
		for {
			if isPrime(gg) {
				fmt.Fprintln(writer, gg)
				break
			}
			gg++
		}
	}
}
