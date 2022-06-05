package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var s float64
	var d float64

	fmt.Println("введите площадь круга")
	fmt.Fscan(os.Stdin, &s)
	d = 2 * math.Sqrt(s/math.Pi)
	l := math.Pi * d
	fmt.Println(d)
	fmt.Println(l)

}
