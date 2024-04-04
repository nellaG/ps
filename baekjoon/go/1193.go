package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var n, dir, sum int
	fmt.Fscanln(reader, &n)
	dir = -1
	if n == 1 {
		fmt.Println("1/1")
		return
	}
	for i := 0; i <= 10000000; i++ {
		if sum >= n {
			sub := sum - n
			s := i - 1 - sub
			if dir > 0 {
				fmt.Printf("%d/%d", s, i-dir*s)
			} else {
				fmt.Printf("%d/%d", i+dir*s, s)
			}
			return

		}
		sum += i
		dir *= -1
	}
}
