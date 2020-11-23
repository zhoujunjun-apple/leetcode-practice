package main

import "testing"

var testCases = []struct {
	arrA []int
	arrB []int
	expL int // the expected max length of common subarray
}{
	{
		arrA: []int{0, 1, 2, 3},
		arrB: []int{1, 2, 4, 5},
		expL: 2,
	},
	{
		arrA: []int{},
		arrB: []int{1},
		expL: 0,
	},
	{
		arrA: []int{0, 1},
		arrB: []int{2, 3},
		expL: 0,
	},
	{
		arrA: []int{1, 2},
		arrB: []int{1, 3, 2, 4},
		expL: 1,
	},
	{
		arrA: []int{1, 2, 3, 4},
		arrB: []int{4, 3, 2, 1},
		expL: 1,
	},
	{
		arrA: []int{1, 2, 1, 3},
		arrB: []int{1, 3, 1, 2},
		expL: 2,
	},
	{
		arrA: []int{1, 2, 3, 4, 5, 6, 7, 8, 0},
		arrB: []int{0, 1, 2, 4, 4, 5, 6, 8, 9},
		expL: 3,
	},
	{
		arrA: []int{1, 2, 3, 4, 5, 6},
		arrB: []int{2, 3, 4, 5, 7, 1},
		expL: 4,
	},
}

func TestRecursionMain(t *testing.T) {
	for i, tc := range testCases {
		got := RecursionMain(tc.arrA, tc.arrB)
		if got != tc.expL {
			t.Errorf("index=%d, expected=%d, got=%d\n", i, tc.expL, got)
		}
	}
}

func TestNativeDP(t *testing.T) {
	for i, tc := range testCases {
		got := NativeDP(tc.arrA, tc.arrB)
		if got != tc.expL {
			t.Errorf("index=%d, expected=%d, got=%d\n", i, tc.expL, got)
		}
	}
}
