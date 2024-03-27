package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)

	fmt.Printf("%d\n%d\n%d\n%d", (a+b)%c, ((a%c)+(b%c))%c, (a*b)%c, ((a%c)*(b%c))%c)
}
