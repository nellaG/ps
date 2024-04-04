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
	var d float64
	d = 2
	for i := 0; i < n; i++ {
		d += math.Pow(2, float64(i))
	}
	fmt.Println(int(math.Pow(float64(d), 2)))

}
