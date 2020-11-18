package main

import (
	"strconv"
)

// BruteForce function is WRONG.
// note: SumOver(start, end, ele) means add ele from start to end.
// The true brute force is to find : given a string s with length N,
// sample out k(1<=k<=N) characters from s without change their relative position.
// when k is 1, there are N subsequences to be checked. each check cost 1. total O(N) time;
// when k is 2, there are SumOver(1, N-1, ele)=N(N-1)/2 subseqs to be checked. each check cost 2. total O(N^2) time;
// when k is 3, there are SumOver(1, N-2, ele(ele+1)/2) subseqs to be checked. each check cost 3. total O(N^3) time;
// so, N + N^2 + N^3 + N^4 + ...
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
// it has the time O(N*N) and space O(N) cost.
func NativeDP(s string) int {
	// dp[i]表示以s[i]结尾的最长递增自序列的长度
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = 1
	}

	ret := 0
	for pos := 0; pos < len(s); pos++ {
		posNum, _ := strconv.Atoi(s[pos : pos+1])
		longest := 0
		// for each pos, look backward ...
		for back := pos - 1; back >= 0; back-- {
			backNum, _ := strconv.Atoi(s[back : back+1])
			// ... look backward to find all possible subsequence ending with s[pos]
			if backNum < posNum && dp[back] > longest {
				// record length of the longer subsequences ending with s[pos]
				longest = dp[back]
			}
		}
		dp[pos] += longest // now dp[pos] is the length of the longest valid subsequence ending with s[pos]

		if dp[pos] > ret { // update the global longest length
			ret = dp[pos]
		}
	}

	return ret
}

func main() {

}
