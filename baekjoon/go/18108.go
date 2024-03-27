package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a int
	var bud int = 2541 - 1998
	fmt.Fscanln(reader, &a)
	fmt.Println(a - bud)
}
