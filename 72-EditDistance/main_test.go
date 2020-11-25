package main

import "testing"

var testCases = []struct {
	s1  string
	s2  string
	ops int
}{
	{
		s1:  "abcd",
		s2:  "ace",
		ops: 2,
	},
	{
		s1:  "abc",
		s2:  "abc",
		ops: 0,
	},
	{
		s1:  "a",
		s2:  "b",
		ops: 1,
	},
	{
		s1:  "a",
		s2:  "ab",
		ops: 1,
	},
	{
		s1:  "a",
		s2:  "bc",
		ops: 2,
	},
	{
		s1:  "abc",
		s2:  "ac",
		ops: 1,
	},

	{
		s1:  "abc",
		s2:  "",
		ops: 3,
	},
	{
		s1:  "aaa",
		s2:  "aac",
		ops: 1,
	},
	{
		s1:  "abcd",
		s2:  "cde",
		ops: 3,
	},
	{
		s1:  "abcde",
		s2:  "ce",
		ops: 3,
	},
}

func TestRecursionMain(t *testing.T) {
	for i, tc := range testCases {
		got := RecursionMain(tc.s1, tc.s2)
		if got != tc.ops {
			t.Errorf("index=%d, s1=%s, s2=%s, expected=%d, got=%d\n", i, tc.s1, tc.s2, tc.ops, got)
		}
	}
}
