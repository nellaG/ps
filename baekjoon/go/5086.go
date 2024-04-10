package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	for {
		fmt.Fscanln(reader, &a, &b)
		if a == 0 && b == 0 {
			break
		}
		if b%a == 0 {
			fmt.Println("factor")
		} else if a%b == 0 {
			fmt.Println("multiple")
		} else {
			fmt.Println("neither")
		}
	}
}
