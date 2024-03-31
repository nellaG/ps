package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var input string
	nmat := make([][]string, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscanln(reader, &input)
		row := make([]string, 15)
		nmat[i] = row
		for j, v := range input {
			nmat[i][j] = string(v)
		}
	}
	for j := 0; j < 15; j++ {
		for i := 0; i < 5; i++ {
			fmt.Fprintf(writer, "%s", nmat[i][j])
		}
	}
}
