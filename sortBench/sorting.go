package sort

import (
	"sort"
)

var sl []int = []int{8, 6, 1, 9, 3, 7, 4, 2, 5}

func SortOne(sl []int) []int {

	for i := 1; i < len(sl); i++ {
		x := sl[i]
		j := i
		for ; j >= 1 && sl[j-1] > x; j-- {
			sl[j] = sl[j-1]
		}
		sl[j] = x
	}
	return sl
}

func SortTwo(sl []int) []int {

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	return sl

}
