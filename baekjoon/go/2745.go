package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	var input string
	var count, digit float64
	fmt.Fscanln(reader, &input, &n)
	l := len(input)
	//strconv.ParseInt()
	for i, v := range input {
		if v <= 57 {
			digit = float64(v - '0')
		} else {
			digit = float64(v - 55)
		}
		count += digit * math.Pow(float64(n), float64(l-i-1))
	}
	fmt.Println(int(count))
}
