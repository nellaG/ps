package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, money int
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		var q, d, ni, p int
		fmt.Fscanln(reader, &money)
		q = money / 25
		money %= 25
		d = money / 10
		money %= 10
		ni = money / 5
		p = money % 5
		fmt.Printf("%d %d %d %d\n", q, d, ni, p)
	}
}
