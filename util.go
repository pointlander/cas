package main

import (
	"log"
	"testing"
)

func BenchmarkCompile(b *testing.B, newCAS func() CAS) {
	algebrite, err := algebriteBundleForBrowserJsBytes()
	if err != nil {
		log.Panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		cas := newCAS()
		b.StartTimer()
		err := cas.Compile(algebrite)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLoad(b *testing.B, newCAS func() CAS) {
	algebrite, err := algebriteBundleForBrowserJsBytes()
	if err != nil {
		log.Panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		cas := newCAS()
		err := cas.Compile(algebrite)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
		err = cas.Load()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRun(b *testing.B, newCAS func() CAS) {
	algebrite, err := algebriteBundleForBrowserJsBytes()
	if err != nil {
		log.Panic(err)
	}

	cas := newCAS()
	err = cas.Compile(algebrite)
	if err != nil {
		b.Fatal(err)
	}
	err = cas.Load()
	if err != nil {
		b.Fatal(err)
	}

	lines := []string{
		"123 + 123",
		"integral(x^2)",
		//"defint(x^2,y,0,sqrt(1-x^2),x,-1,1)",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, line := range lines {
			_, err = cas.Run(line)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
