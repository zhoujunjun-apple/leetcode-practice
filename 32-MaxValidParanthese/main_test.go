package main

import (
	"testing"
)

func Test_DPMaxValidParenthese(t *testing.T) {
	ts := []struct {
		input string
		exput int
	}{
		{
			input: ")())()())((())))",
			exput: 6,
		},
		{
			input: ")())()())",
			exput: 4,
		},
		{
			input: "(()())",
			exput: 6,
		},
		{
			input: "(()(((()()",
			exput: 4,
		},
		{
			input: ")())()()",
			exput: 4,
		},
		{
			input: "((",
			exput: 0,
		},
		{
			input: "))",
			exput: 0,
		},
		{
			input: "(()",
			exput: 2,
		},
		{
			input: "))()(",
			exput: 2,
		},
	}

	for _, tc := range ts {
		got := DPMaxValidParenthese(tc.input)
		if got != tc.exput {
			t.Errorf("input=%s, expected=%d, got=%d", tc.input, tc.exput, got)
		}
	}
}
