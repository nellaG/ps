package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k, count int
	fmt.Fscanln(reader, &n, &k)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
		if count == k {
			fmt.Println(i)
			return
		}
	}
	if count < k {
		fmt.Println(0)
	}
}
