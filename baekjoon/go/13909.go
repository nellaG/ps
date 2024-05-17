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
	fmt.Fscanln(reader, &n)
	fmt.Println(int(math.Sqrt(float64(n))))
}
