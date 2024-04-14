package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, count int
	fmt.Fscan(reader, &n)
	for i := n - 2; i > 0; i-- {
		count += ((i + 1) * i / 2)
	}
	fmt.Println(count)
	fmt.Println(3)
}
