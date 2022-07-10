package sort

import "testing"

func BenchmarkSortOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortOne(sl)
	}
}

func BenchmarkSortTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortTwo(sl)
	}
}
