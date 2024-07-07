package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	fmt.Println(square(n))
}

func concat(s string, punch bool) string {
	lines := strings.Split(s, "\n")
	hz := ""
	for _, line := range lines {
		var pl = strings.Repeat(" ", len(line))
		if !punch {
			pl = line
		}
		hz += line + pl + line + "\n"
	}
	return strings.TrimRight(hz, "\n")
}

func square(n int) string {
	var buf bytes.Buffer
	if n <= 3 {
		buf.WriteString("***\n* *\n***")
		return buf.String()
	}
	subSq := square(n / 3)
	buf.WriteString(concat(subSq, false) + "\n" + concat(subSq, true) + "\n" + concat(subSq, false))

	return buf.String()

}
