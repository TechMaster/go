package main

import "testing"

/*
Hướng dẫn kiểm tra tốc độ thực thi hàm
go test -bench .
*/
func Benchmark_RemoveItemSliceNotKeepOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeItemSliceNotKeepOrder([]string{"a", "b", "c", "d", "e", "f"}, 2)
	}
}

func Benchmark_RemoveItemSliceNotKeepOrder2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeItemSliceNotKeepOrder2([]string{"a", "b", "c", "d", "e", "f"}, 2)
	}
}

func Benchmark_RemoveItemSliceKeepOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeItemSliceKeepOrder([]string{"a", "b", "c", "d", "e", "f"}, 2)
	}
}

func Benchmark_RemoveItemSliceKeepOrder2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeItemSliceKeepOrder2([]string{"a", "b", "c", "d", "e", "f"}, 2)
	}
}
