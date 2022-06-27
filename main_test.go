package main

import "testing"

func BenchmarkTest2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Test2()
	}
}
