package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, x, y, gcd_, top, bottom int
	fmt.Fscanln(reader, &a, &x)
	fmt.Fscanln(reader, &b, &y)
	top = a*y + b*x
	bottom = x * y
	gcd_ = gcd(top, bottom)
	fmt.Println(top/gcd_, bottom/gcd_)
}
