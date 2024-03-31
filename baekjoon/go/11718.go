package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprint(writer, text)
	}
}
