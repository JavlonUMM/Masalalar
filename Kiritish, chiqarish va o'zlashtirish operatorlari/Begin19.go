package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Print("x1: ")
	fmt.Scan(&x1)
	fmt.Print("y1: ")
	fmt.Scan(&y1)
	fmt.Print("x2: ")
	fmt.Scan(&x2)
	fmt.Print("y2: ")
	fmt.Scan(&y2)

	uzunlik1 := math.Abs(x2 - x1)
	uzunlik2 := math.Abs(y2 - y1)

	perimeteri := 2 * (uzunlik1 + uzunlik2)

	yuzasi := uzunlik1 * uzunlik2

	fmt.Printf("Perimetri: %.2f\n", perimeteri)
	fmt.Printf("Yuzasi: %.2f\n", yuzasi)
}
