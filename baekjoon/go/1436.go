package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, count int
	fmt.Fscan(reader, &n)
	for i := 1; i <= 100000000; i++ {
		if strings.Contains(strconv.Itoa(i), "666") {
			count++
			if n == count {
				fmt.Println(i)
				return
			}
		}
	}
}
