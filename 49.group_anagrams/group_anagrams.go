// https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
// 49. 字母异位词分组
//
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

package groupanagrams

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)
	res := [][]string{}

	for _, str := range strs {
		k := getKey(str)
		if i, ok := m[k]; ok {
			res[i] = append(res[i], str)
		} else {
			res = append(res, []string{str})
			m[k] = len(res) - 1
		}
	}

	return res
}

func getKey(str string) string {
	b := []byte(str)
	slices.Sort(b)
	return string(b)
}
