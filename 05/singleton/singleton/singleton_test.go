package singleton

import (
	"testing"
)

func Benchmark_GetInstance_01(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetInstanceFunc()
	}
}

func Benchmark_GetInstance_02(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetInstanceSync()
	}
}

func Benchmark_GetInstance_03(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetInstanceMutex()
	}
}

func Benchmark_GetInstance_04(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetInstanceInit()
	}
}