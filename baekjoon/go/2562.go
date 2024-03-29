package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var num int
	var idx int
	max_ := 1
	for i := 0; i < 9; i++ {
		fmt.Fscan(reader, &num)
		if num > max_ {
			max_ = num
			idx = i + 1
		}
	}
	fmt.Printf("%d\n%d", max_, idx)
}
