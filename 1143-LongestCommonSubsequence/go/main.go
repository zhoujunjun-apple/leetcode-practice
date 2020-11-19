package main

// LongestCommonSubsequenceDP function use dp
func LongestCommonSubsequenceDP(one, two string) int {
	oneLen, twoLen := len(one), len(two)
	if oneLen < 1 || twoLen < 1 {
		return 0
	}

	// dp[i][j]表示one[0:i]和two[0:j]子字符串之间的最长公共子序列的长度
	dp := make([][]int, oneLen)
	for i := range dp {
		dp[i] = make([]int, twoLen)
	}

	// initialize the first column dp[j][0]
	for i := 0; i < oneLen; i++ {
		if one[i] == two[0] {
			for j := i; j < oneLen; j++ {
				dp[j][0] = 1
			}
		}
	}

	// initialize the first row dp[0][j]
	for i := 0; i < twoLen; i++ {
		if two[i] == one[0] {
			for j := i; j < twoLen; j++ {
				dp[0][j] = 1
			}
		}
	}

	// fill up the dp table
	for oneIdx := 1; oneIdx < oneLen; oneIdx++ {
		for twoIdx := 1; twoIdx < twoLen; twoIdx++ {
			if one[oneIdx] == two[twoIdx] {
				// 如果one[i]和two[j]相同，则one[i]和two[j]必定是one[0:i]和two[0:j]子字符串之间公共字序列的末尾字符
				dp[oneIdx][twoIdx] = dp[oneIdx-1][twoIdx-1] + 1
			} else {
				// 否则，从dp[i-1][j-1], dp[i-1][j]和dp[i][j-1]之间选择最大的，
				// 又因为dp[i-1][j]和dp[i][j-1]必然不小于dp[i-1][j-1]，因此
				// 只需要从dp[i-1][j]和dp[i][j-1]中选择较大值即可
				if dp[oneIdx-1][twoIdx] > dp[oneIdx][twoIdx-1] {
					dp[oneIdx][twoIdx] = dp[oneIdx-1][twoIdx]
				} else {
					dp[oneIdx][twoIdx] = dp[oneIdx][twoIdx-1]
				}
			}
		}
	}

	return dp[oneLen-1][twoLen-1]
}

func main() {

}
