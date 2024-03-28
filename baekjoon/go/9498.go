package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a int
	fmt.Fscanln(reader, &a)
	switch {
	case 90 <= a && a <= 100:
		fmt.Println("A")
	case 80 <= a && a < 90:
		fmt.Println("B")
	case 70 <= a && a < 80:
		fmt.Println("C")
	case 60 <= a && a < 70:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}
