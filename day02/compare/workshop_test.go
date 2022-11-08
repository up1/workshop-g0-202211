package compare_test

import (
	"day02/compare"
	"testing"
)

func BenchmarkV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compare.V1()
    }
}