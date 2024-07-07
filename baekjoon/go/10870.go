package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-2) + fibo(n-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	fmt.Println(fibo(n))
}
