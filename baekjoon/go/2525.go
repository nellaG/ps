package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscanln(reader, &a, &b)
	fmt.Fscanln(reader, &c)
	hour := c / 60
	minute := c % 60
	a += hour
	b += minute
	if b >= 60 {
		a += 1
		b -= 60
	}
	fmt.Printf("%d %d\n", a%24, b)
}
