package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a string
	fmt.Fscanln(reader, &a)
	var count int

	for _, v := range a {
		asc := v - 65
		switch {
		case 0 <= asc && asc < 3:
			count += 3
		case 3 <= asc && asc < 6:
			count += 4
		case 6 <= asc && asc < 9:
			count += 5
		case 9 <= asc && asc < 12:
			count += 6
		case 12 <= asc && asc < 15:
			count += 7
		case 15 <= asc && asc < 19:
			count += 8
		case 19 <= asc && asc < 22:
			count += 9
		case 22 <= asc && asc < 26:
			count += 10
		}
	}
	fmt.Println(count)
}
