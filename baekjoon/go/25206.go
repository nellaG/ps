package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var d, l string
	var grade, sum, gsum float64
	table := map[string]float64{"A+": 4.5, "A0": 4.0, "B+": 3.5, "B0": 3.0,
		"C+": 2.5, "C0": 2.0, "D+": 1.5, "D0": 1.0, "F": 0.0}

	for i := 0; i < 20; i++ {
		fmt.Fscanln(reader, &d, &grade, &l)
		if l == "P" {
			continue
		}
		gsum += grade * table[l]
		sum += grade
	}
	fmt.Println(gsum / float64(sum))
}
