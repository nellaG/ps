package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input, n, digit int
	var ans string
	fmt.Fscanln(reader, &input, &n)
	//strconv.ParseInt()

	for {
		v := input % n
		if v < 10 {
			digit = 48
		} else {
			digit = 55
		}
		if input < n {
			ans = string(v+digit) + ans
			break
		}
		input /= n
		ans = string(v+digit) + ans
	}
	fmt.Println(ans)
}
