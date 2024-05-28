package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k, top int
	top = 1
	fmt.Fscanln(reader, &n, &k)
	for i := n; i > n-k; i-- {
		top *= i
	}
	for i := 1; i <= k; i++ {
		top /= i
	}
	fmt.Println(top)
}
