package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	for j := 0; j <= n/3; j++ {
		for i := 0; i <= n/5; i++ {
			if i*5+j*3 == n {
				fmt.Println(i + j)
				return
			}
		}
	}
	fmt.Println(-1)
}
