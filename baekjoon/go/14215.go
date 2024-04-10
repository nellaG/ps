package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)
	if a+c <= b {
		fmt.Println(a + c + (a + c - 1))
	} else if b+c <= a {
		fmt.Println(b + c + (b + c - 1))
	} else if a+b <= c {
		fmt.Println(b + a + (b + a - 1))
	} else {
		fmt.Println(b + a + c)
	}
}
