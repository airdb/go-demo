package no6_strings_builder

import (
	"strings"
	"testing"
)

var str string

var strs = []string{
	"here's",
	"a",
	"some",
	"long",
	"list",
	"of",
	"strings",
	"for",
	"you",
}

func BuildStrRaw(strs []string) string {
	var s string

	for _, v := range strs {
		s += v
	}

	return s
}

func BuildStrBuilder(strs []string) string {
	b := strings.Builder{}

	for _, v := range strs {
		b.WriteString(v)
	}

	return b.String()
}

func BuildStrPreAllocBuilder(strs []string) string {
	b := strings.Builder{}
	b.Grow(128)

	for _, v := range strs {
		b.WriteString(v)
	}

	return b.String()
}

func BenchmarkStringBuildRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = BuildStrRaw(strs)
	}
}

func BenchmarkStringBuildBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = BuildStrBuilder(strs)
	}
}

func BenchmarkStringPreAllocBuildBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = BuildStrPreAllocBuilder(strs)
	}
}
