package main

import "testing"

func TestBruteForce(t *testing.T) {
	ts := []struct {
		input string
		exput int
	}{
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

	for i, tc := range ts {
		got := BruteForce(tc.input)
		if got != tc.exput {
			t.Errorf("%d-th input=%s, expected=%d, got=%d", i, tc.input, tc.exput, got)
		}
	}
}
