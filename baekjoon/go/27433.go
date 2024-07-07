package main

import (
	"bufio"
	"fmt"
	"os"
)

func fact(n int) int {
	if n <= 1 {
		return n
	}
	return n * fact(n-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	fmt.Println(fact(n))
}
