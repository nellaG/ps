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
	var n int
	fmt.Fscanln(reader, &n)
	if n == 1 {
		return
	}
	for !isPrime(n) {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				n /= i
				fmt.Fprintln(writer, i)
				break
			}
		}
	}
	fmt.Fprintln(writer, n)
}
