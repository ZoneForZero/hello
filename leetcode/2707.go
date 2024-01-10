package lc

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	//给你一个下标从 0 开始的字符串 s 和一个单词字典 dictionary 。你需要将 s 分割成若干个 互不重叠 的子字符串，每个子字符串都在 dictionary 中出现过。s 中可能会有一些 额外的字符 不在任何子字符串中。
	//
	//请你采取最优策略分割 s ，使剩下的字符 最少 。
	mapDict := map[string]bool{}
	for _, v := range dictionary {
		mapDict[v] = true
	}
	lenS := len(s)
	dp := make([]int, lenS+1)
	dp[0] = 0
	for i := 1; i <= lenS; i++ {
		dp[i] = 51
	}
	for endIndex := 1; endIndex <= lenS; endIndex++ {
		tempStr := ""
		for start := endIndex - 1; start >= 0; start-- {
			tempStr = string(s[start]) + tempStr
			lenT := len(tempStr)
			dp[endIndex] = min(dp[endIndex], dp[endIndex-lenT]+lenT)
			if mapDict[tempStr] {
				fmt.Println("看看2222", tempStr, dp[endIndex], dp[endIndex-lenT], lenT)
				dp[endIndex] = min(dp[endIndex], dp[endIndex-lenT])
			}
		}
	}
	fmt.Println(dp)
	return dp[lenS]
}
func init() {
	fmt.Println("???")
	fmt.Println(minExtraChar("leetcode", []string{"leet", "code", "leetcode"}))
}
