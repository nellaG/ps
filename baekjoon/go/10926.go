package main

import (
	"bufio"
	"fmt"
	"os"
)

func ps10926(a string) string {
	return a + "??!"

}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a string
	fmt.Fscanln(reader, &a)
	fmt.Println(ps10926(a))
}
