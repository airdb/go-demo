package no2_benchstat

import (
	"strconv"
	"testing"
)

func MyItoa(i int) string {
	return strconv.Itoa(i)
	//return fmt.Sprint(i)
}

var r string

func BenchmarkMyItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = MyItoa(i)
	}
}
