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

func main() {

}
