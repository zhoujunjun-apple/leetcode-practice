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

func NativeDP(s string) int {
	dp := make([]int, len(s))

}

func main() {

}
