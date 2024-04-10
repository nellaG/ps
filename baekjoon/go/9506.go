package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, sum_ int
	var str string
	for {
		fmt.Fscanln(reader, &n)
		str = "1"
		sum_ = 1
		if n == -1 {
			break
		}
		for i := 2; i < n; i++ {
			if n%i == 0 {
				sum_ += i
				str += fmt.Sprintf(" + %d", i)
			}
		}
		if sum_ == n {
			fmt.Printf("%d = %s\n", n, str)
		} else {
			fmt.Printf("%d is NOT perfect.\n", n)
		}

	}
}
