package test

import (
	"fmt"
	"testing"
)

var r string

func BenchmarkFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = fmt.Sprint(i)
	}
}
