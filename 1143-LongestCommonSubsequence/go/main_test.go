package main

import "testing"

var ts = []struct {
	one    string
	two    string
	common int
}{
	{
		one:    "1",
		two:    "",
		common: 0,
	},
	{
		one:    "2",
		two:    "2",
		common: 1,
	},
	{
		one:    "12",
		two:    "34",
		common: 0,
	},
	{
		one:    "12345678",
		two:    "135990",
		common: 3,
	},
	{
		one:    "12345",
		two:    "54321",
		common: 1,
	},
	{
		one:    "123454321",
		two:    "0011220395",
		common: 4,
	},
	{
		one:    "1358924680",
		two:    "1358924680",
		common: 10,
	},
	{
		one:    "1111",
		two:    "1213141516",
		common: 4,
	},
}

func TestLongestCommonSubsequenceDP(t *testing.T) {
	for i, tt := range ts {
		got := LongestCommonSubsequenceDP(tt.one, tt.two)
		if got != tt.common {
			t.Errorf("index=%d: one=%s, two=%s, expected=%d, got=%d", i, tt.one, tt.two, tt.common, got)
		}
	}
}
