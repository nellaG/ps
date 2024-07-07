package main

import (
	"bufio"
	"fmt"
	"os"
)

func countRec(s string, l int, r int, c int) (int, int) {
	if l >= r {
		return 1, c
	}
	if s[l] != s[r] {
		return 0, c
	}
	return countRec(s, l+1, r-1, c+1)
}

func isPal(s string) (int, int) {
	return countRec(s, 0, len(s)-1, 1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	var s string
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)
		fmt.Println(isPal(s))
	}
}
