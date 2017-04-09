package DP2

func LongestCommonSubsequence(s1 string, s2 string, len1 int, len2 int) int {
	if len1 == -1 || len2 == -1 {
		return 0
	} else if s1[len1] == s2[len2] {
		return 1 + LongestCommonSubsequence(s1, s2, len1 - 1, len2 - 1)
	} else {
		return max(
			LongestCommonSubsequence(s1, s2, len1 - 1, len2),
			LongestCommonSubsequence(s1, s2, len1, len2 - 1))
	}
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
