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
	if a+b+c != 180 {
		fmt.Println("Error")
		return
	}
	if a == b || b == c || a == c {
		if a == b && b == c {
			fmt.Println("Equilateral")
		} else {
			fmt.Println("Isosceles")
		}
		return
	}
	fmt.Println("Scalene")
}
