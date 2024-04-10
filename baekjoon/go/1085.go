package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, w, h int
	fmt.Scanf("%d %d %d %d", &x, &y, &w, &h)
	fmt.Println(min(x, y, int(math.Min(float64(w-x), float64(h-y)))))

}
