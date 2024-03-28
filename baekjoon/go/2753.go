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
	if a%4 == 0 && a%100 != 0 || a%400 == 0 {
		fmt.Println(1)
		return
	}
	fmt.Println(0)
}
