package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var total, num, price, count int
	fmt.Fscanln(reader, &total)
	fmt.Fscanln(reader, &num)

	for i := 0; i < num; i++ {
		fmt.Fscanln(reader, &price, &count)
		total -= (price * count)
	}
	if total == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
