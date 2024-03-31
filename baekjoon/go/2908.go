package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var a, b string
	fmt.Fscanln(reader, &a, &b)
	var numA, numB int
	for i := 0; i < 3; i++ {
		valA, _ := strconv.Atoi(string(a[i]))
		numA += valA * int(math.Pow10(i))
		valB, _ := strconv.Atoi(string(b[i]))
		numB += valB * int(math.Pow10(i))
	}
	if numA > numB {
		fmt.Println(numA)
	} else {
		fmt.Println(numB)
	}
}
