package main

// NativeDP function represents the native dp implementation
func NativeDP(a, b []int) int {
	aLen, bLen := len(a), len(b)
	if aLen == 0 || bLen == 0 {
		return 0
	}

	dp := make([][]int, aLen)
	for i := range dp {
		dp[i] = make([]int, bLen)
	}

	for ai := 0; ai < aLen; ai++ {
		if a[ai] == b[0] {
			dp[ai][0] = 1
		}
	}

	for bi := 0; bi < bLen; bi++ {
		if b[bi] == a[0] {
			dp[0][bi] = 1
		}
	}

	mr := 0
	for ai := 1; ai < aLen; ai++ {
		for bi := 1; bi < bLen; bi++ {
			if a[ai] == b[bi] {
				c := dp[ai-1][bi-1] + 1
				dp[ai][bi] = c
				if c > mr {
					mr = c
				}
			}
		}
	}

	return mr
}

// RecursionMain function represents the entry of recursion implementation
func RecursionMain(a, b []int) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	ret := 0
	for ai := 0; ai < len(a); ai++ {
		for bi := 0; bi < len(b); bi++ {
			r := PureRecursion(a, b, ai, bi)
			if r > ret {
				ret = r
			}
		}
	}
	return ret
}

// PureRecursion function represents the maximum length of common subarray
// ending with a[ai] and b[bi]
func PureRecursion(a, b []int, ai, bi int) int {
	if ai == 0 || bi == 0 {
		if a[ai] == b[bi] {
			return 1
		}
		return 0
	}

	if a[ai] == b[bi] {
		return PureRecursion(a, b, ai-1, bi-1) + 1
	}
	return 0
}

func main() {}
