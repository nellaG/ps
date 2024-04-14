package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a1, a0, c, n0 int
	fmt.Fscan(reader, &a1, &a0, &c, &n0)
	for i := n0; i < 10000; i++ {
		x := a1*i + a0
		if x > c*i {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(1)
}
