package main

// RecursionMain function is the entry of recursion method
func RecursionMain(s1, s2 string) int {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len == 0 {
		return s2Len
	}
	if s2Len == 0 {
		return s1Len
	}

	return Recursion(s1, s2, s1Len-1, s2Len-1)
}

// Recursion function Recursion(s1, s2, i, j)表示将s1[0, i]编辑为s2[0, j]所需的最少操作步数ops
func Recursion(s1, s2 string, i, j int) int {
	ops := 0
	if i == 0 {
		for k := 0; k <= j; k++ {
			if s1[0] == s2[k] {
				return j
			}
		}
		return j + 1
	}
	if j == 0 {
		for k := 0; k <= i; k++ {
			if s2[0] == s1[k] {
				return i
			}
		}
		return i + 1
	}

	// Recursion(s1, s2, i-1, j) + 1: 前者表示s1[0,i-1]编辑为s2[0,j]的ops, +1表示对s1[i]的删除操作;
	// Recursion(s1, s2, i, j-1) + 1: 前者表示s1[0,i]编辑为s2[0,j-1]的ops，+1表示s2[j]的插入操作;
	// Recursion(s1, s2, i-1, j-1): 表示s1[0,i-1]编辑为s2[0,j-1]的ops，当s1[i]不等于s2[j]时，加上将s1[i]替换为s2[j]的操作;
	if s1[i] == s2[j] {
		ops = getMin(Recursion(s1, s2, i-1, j)+1, Recursion(s1, s2, i, j-1)+1, Recursion(s1, s2, i-1, j-1))
	} else {
		ops = getMin(Recursion(s1, s2, i-1, j)+1, Recursion(s1, s2, i, j-1)+1, Recursion(s1, s2, i-1, j-1)+1)
	}
	return ops
}

func getMin(i, j, k int) int {
	min := i
	if i < j {
		if i < k {
			min = i
		} else {
			min = k
		}
	} else {
		if j < k {
			min = j
		} else {
			min = k
		}
	}
	return min
}

// NativeDP function implement dp method rather than recursion
func NativeDP(s1, s2 string) int {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len == 0 {
		return s2Len
	}
	if s2Len == 0 {
		return s1Len
	}

	dp := make([][]int, s1Len)
	for i := range dp {
		dp[i] = make([]int, s2Len)
	}

	// fill up the first column
	for ri := 0; ri < s1Len; ri++ {
		if s1[ri] == s2[0] {
			for k := ri; k < s1Len; k++ {
				dp[k][0] = k
			}
			break
		} else {
			dp[ri][0] = ri + 1
		}
	}

	// fill up the first row
	for ci := 0; ci < s2Len; ci++ {
		if s2[ci] == s1[0] {
			for k := ci; k < s2Len; k++ {
				dp[0][k] = k
			}
			break
		} else {
			dp[0][ci] = ci + 1
		}
	}

	// same logic with recursion
	for ri := 1; ri < s1Len; ri++ {
		for ci := 1; ci < s2Len; ci++ {
			if s1[ri] == s2[ci] {
				dp[ri][ci] = getMin(dp[ri-1][ci]+1, dp[ri][ci-1]+1, dp[ri-1][ci-1])
			} else {
				dp[ri][ci] = getMin(dp[ri-1][ci], dp[ri][ci-1], dp[ri-1][ci-1]) + 1
			}
		}
	}

	return dp[s1Len-1][s2Len-1]
}

func main() {

}
