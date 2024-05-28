package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, c int
	c = 1
	fmt.Fscanln(reader, &n)
	for i := 1; i <= n; i++ {
		c *= i
	}
	fmt.Println(c)
}
