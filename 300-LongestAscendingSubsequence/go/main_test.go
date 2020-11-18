package main

import "testing"

var ts = []struct {
	input string
	exput int
}{
	{
		input: "1252345",
		exput: 5,
	},
	{
		input: "43210",
		exput: 1,
	},
	{
		input: "21340532",
		exput: 4,
	},
	{
		input: "012345000",
		exput: 6,
	},
}

func TestBruteForce(t *testing.T) {
	for i, tc := range ts {
		got := BruteForce(tc.input)
		if got != tc.exput {
			t.Errorf("%d-th input=%s, expected=%d, got=%d", i, tc.input, tc.exput, got)
		}
	}
}

func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "21340532"
		BruteForce(s)
	}
}

func TestNativeDP(t *testing.T) {
	for i, tc := range ts {
		got := NativeDP(tc.input)
		if got != tc.exput {
			t.Errorf("%d-th input=%s, expected=%d, got=%d", i, tc.input, tc.exput, got)
		}
	}
}

func BenchmarkNativeDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "21340532"
		NativeDP(s)
	}
}
