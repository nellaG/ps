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
	var x, c int
	for {
		c = 0
		fmt.Fscanln(reader, &x)
		if x == 0 {
			break
		}
		for i := x + 1; i <= 2*x; i++ {
			if isPrime(i) {
				c++
			}
		}
		fmt.Println(c)
	}
}
