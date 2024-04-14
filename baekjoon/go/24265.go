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
	for i := 1; i < n; i++ {
		count += i
	}
	fmt.Println(count)
	fmt.Println(2)
}
