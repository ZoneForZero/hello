package lc

import "sort"

type ac struct {
	strIndex int
	from     string
	to       string
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	lenS := len(s)
	k := len(indices)
	acs := make([]ac, 0, k)
	for i := 0; i < k; i++ {
		acs = append(acs, ac{
			indices[i],
			sources[i],
			targets[i],
		})
	}
	sort.Slice(acs, func(i, j int) bool {
		return acs[i].strIndex < acs[j].strIndex
	})
	ans := ""
	// 闭区间
	strEndIndex := 0
	for i := 0; i < k; i++ {
		strindex, fromStr, toStr := acs[i].strIndex, acs[i].from, acs[i].to
		// pre handler
		if strindex > strEndIndex {
			ans += s[strEndIndex:strindex]
			strEndIndex = strindex
		}
		lenP := len(fromStr)
		if strindex+lenP <= lenS && s[strindex:strindex+lenP] == fromStr {
			ans += toStr
			strEndIndex = strindex + lenP
		}
	}
	ans += s[strEndIndex:]
	return ans
}
