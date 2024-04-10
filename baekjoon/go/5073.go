package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c int
	for {
		fmt.Fscan(reader, &a, &b, &c)
		if a == 0 && b == 0 && c == 0 {
			break
		}
		if a+c <= b || b+c <= a || a+b <= c {
			fmt.Println("Invalid")
		} else if a == b || b == c || a == c {
			if a == b && b == c {
				fmt.Println("Equilateral")
			} else {
				fmt.Println("Isosceles")
			}
		} else {
			fmt.Println("Scalene")
		}
	}
}
