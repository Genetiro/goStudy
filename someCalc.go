package main

import (
	"fmt"
	"os"
)

func main() {
	var a int

	fmt.Println("введите трехзначное число")
	fmt.Fscan(os.Stdin, &a)
	b := a % 10
	c := (a % 100) / 10
	d := a / 100
	fmt.Println(d, " сотни")
	fmt.Println(c, " десятка")
	fmt.Println(b, " единиц")
}
