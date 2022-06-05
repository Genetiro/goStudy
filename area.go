package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	var b int

	fmt.Println("введите длину прямоугольника")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println("введите ширину прямоугольника")
	fmt.Fscan(os.Stdin, &b)

	fmt.Println("Площадь прямоугольника равна ", a*b)
}
