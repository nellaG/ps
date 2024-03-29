package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var num int
	fmt.Fscanln(reader, &num)
	var score, sum, max float64
	for i := 0; i < num; i++ {
		fmt.Fscan(reader, &score)
		if score > max {
			max = score
		}
		sum += (score * 100)
	}
	fmt.Printf("%f\n", (sum / (float64(num) * max)))
}
