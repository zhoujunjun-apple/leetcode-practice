package main

import (
	"strconv"
)

// BruteForce function check each position s[i].
// start at s[i], find the longest ascending subsequece
func BruteForce(s string) int {
	snum := make([]int, len(s))
	for i, c := range s {
		snum[i], _ = strconv.Atoi(string(c))
	}

	ret := 0
	for i := 0; i < len(s); i++ {
		cur := snum[i]
		length := 1
		for start := i + 1; start < len(s); start++ {
			cmp := snum[start]
			if cmp > cur {
				cur = cmp
				length++
			}
		}

		if length > ret {
			ret = length
		}
	}
	return ret
}

// NativeDP function implement native dp solution
// it has the same time and space cost with brute force solution
func NativeDP(s string) int {
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = 1
	}

	ret := 0
	for pos := 0; pos < len(s); pos++ {
		posNum, _ := strconv.Atoi(s[pos : pos+1])
		longest := 0
		for back := pos - 1; back >= 0; back-- {
			backNum, _ := strconv.Atoi(s[back : back+1])
			if backNum < posNum && dp[back] > longest {
				longest = dp[back]
			}
		}
		dp[pos] += longest

		if dp[pos] > ret {
			ret = dp[pos]
		}
	}

	return ret
}

func main() {

}
