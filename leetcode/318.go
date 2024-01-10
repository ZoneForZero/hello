package lc

import (
	"acm/tools"
	"fmt"
	"sort"
)

// 给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j]) 的最大值，
// 并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。
func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	lenW := len(words)
	ws := make([]int32, lenW)
	ans := 0
	for i := 0; i < lenW; i++ {
		str := words[i]
		if ws[i] == 0 {
			for _, v := range str {
				ws[i] = ws[i] | int32(tools.Pow2(int(v-'a'+1)))
			}
		}
		for j := i + 1; j < lenW; j++ {
			if ws[j] == 0 {
				for _, v := range words[j] {
					ws[j] = ws[j] | int32(tools.Pow2(int(v-'a'+1)))
				}
			}
			if ws[i]&ws[j] == 0 {
				lenW = j
				ta := len(str) * len(words[j])
				if ta > ans {
					ans = ta
				}
				break
			}
		}
	}
	return ans
}
func init() {
	strs := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	fmt.Println(maxProduct(strs))
}
