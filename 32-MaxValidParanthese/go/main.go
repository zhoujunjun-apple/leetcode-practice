package main

func DPMaxValidParenthese(s string) int {
	slen := len(s)

	// dp[i]表示以s[i]为结尾的最长有效括号字串的长度
	dp := make([]int, slen) // golang initialize all the element to zero
	for idx := 0; idx < slen; idx++ {
		if s[idx] == ')' {
			// if current char is '(', dp[idx] must be 0
			bIdx := idx - 1
			if bIdx >= 0 {
				// make sure index is valid
				if s[bIdx] == '(' {
					// if the previous is '(' and the current is ')',
					// then dp[idx] must be at least 2
					dp[idx] = dp[idx] + 2

					// try to combine the previous valid substring
					bIdx = bIdx - 1
					if bIdx >= 0 {
						dp[idx] = dp[idx] + dp[bIdx]
					}
				} else if bVal := dp[bIdx]; bVal > 0 {
					// the previous is ')', jump over the valid substring
					bIdx = bIdx - dp[bIdx]
					if bIdx >= 0 && s[bIdx] == '(' {
						// look like this:
						// bIdx  ............   idx
						//  (   valid substring  )
						// this can combine into a new valid substring
						dp[idx] = dp[idx] + bVal + 2

						// try to combine the previous valid substring
						bIdx = bIdx - 1
						if bIdx >= 0 {
							dp[idx] = dp[idx] + dp[bIdx]
						}
					}
				}
			}
		}
	}

	// find the maxmium length of valid parenthese substring
	ret := 0
	for _, v := range dp {
		if v > ret {
			ret = v
		}
	}

	return ret
}

func main() {

}
