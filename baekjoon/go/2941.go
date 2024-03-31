package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var input string
	fmt.Fscanln(reader, &input)
	cro := []string{"dz=", "c=", "c-", "d-", "lj", "nj", "s=", "z="}
	for _, v := range cro {
		input = strings.Replace(input, v, "1", -1)

	}

	fmt.Fprint(writer, len(input))
}
