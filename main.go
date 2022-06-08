package main

import (
	"fmt"
)

func main() {
	var sl []int = []int{8, 6, 1, 9, 3, 7, 4, 2, 5}
	for i := 1; i < len(sl); i++ {
		x := sl[i]
		j := i
		for ; j >= 1 && sl[j-1] > x; j-- {
			sl[j] = sl[j-1]
		}
		sl[j] = x
	}
	fmt.Println(sl)

}
