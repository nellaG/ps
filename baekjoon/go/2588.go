package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	var result int
	var subResult int
	for i := 0; i < 3; i++ {
		digit := (b % int(math.Pow(10, float64(i+1))))
		subResult = a * digit
		result += subResult
		fmt.Println(subResult / int(math.Pow(10, float64(i))))
		b -= digit
	}
	fmt.Println(result)
}
