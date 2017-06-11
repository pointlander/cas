package main

import "testing"

func BenchmarkOttoInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewOtto()
	}
}

func BenchmarkOttoCompile(b *testing.B) {
	BenchmarkCompile(b, NewOtto)
}

func BenchmarkOttoLoad(b *testing.B) {
	BenchmarkLoad(b, NewOtto)
}

func BenchmarkOttoRun(b *testing.B) {
	BenchmarkRun(b, NewOtto)
}
