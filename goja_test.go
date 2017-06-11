package main

import "testing"

func BenchmarkGOJAInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewGOJA()
	}
}

func BenchmarkGOJACompile(b *testing.B) {
	BenchmarkCompile(b, NewGOJA)
}

func BenchmarkGOJALoad(b *testing.B) {
	BenchmarkLoad(b, NewGOJA)
}

func BenchmarkGOJARun(b *testing.B) {
	BenchmarkRun(b, NewGOJA)
}
