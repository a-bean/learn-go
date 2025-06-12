package main

import (
	"fmt"
	"strconv"
)

// 03 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	// 使用map记录字符最后出现的位置
	lastOccurred := make(map[rune]int)
	start, maxLength := 0, 0

	for i, ch := range s {
		// 如果字符已经在当前窗口中出现过，更新起始位置
		if lastPos, ok := lastOccurred[ch]; ok && lastPos >= start {
			start = lastPos + 1
		}
		// 更新最大长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		// 记录字符位置
		lastOccurred[ch] = i
	}

	return maxLength
}

// 30 串联所有单词的子串 https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
// findSubstring 查找字符串中所有单词串联形成的子串的起始位置
// s: 待搜索的字符串
// words: 要串联的单词数组
// 返回所有符合条件的子串的起始索引
func findSubstring(s string, words []string) []int {
	// 处理边界情况
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount

	// 如果字符串长度小于所需总长度，直接返回
	if len(s) < totalLen {
		return []int{}
	}

	// 初始化单词频率表
	wordFreq := make(map[string]int, wordCount)
	for _, word := range words {
		wordFreq[word]++
	}

	result := make([]int, 0)

	// 考虑所有可能的起始位置
	for start := 0; start <= len(s)-totalLen; start++ {
		seen := make(map[string]int)
		matched := 0

		// 从当前位置开始，每次截取一个单词长度进行匹配
		for i := 0; i < wordCount; i++ {
			pos := start + i*wordLen
			currWord := s[pos : pos+wordLen]

			// 如果当前单词不在词表中，直接跳出
			expectedCount, exists := wordFreq[currWord]
			if !exists {
				break
			}

			// 统计当前单词出现次数
			seen[currWord]++

			// 如果当前单词出现次数超过预期，跳出
			if seen[currWord] > expectedCount {
				break
			}

			matched++
		}

		// 如果所有单词都匹配成功，记录起始位置
		if matched == wordCount {
			result = append(result, start)
		}
	}

	return result
}

// 优化:
// 减少了重复计算
// 避免了频繁的 map 创建和销毁
// 利用了单词长度固定的特性
// 空间复杂度保持在 O(K)，其中 K 是不同单词的数量
// 实际运行时间显著降低，特别是对于长字符串和大量单词的情况
func findSubstring1(s string, words []string) []int {
	// 处理边界情况
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordCount := len(words)
	strLen := len(s)
	totalLen := wordLen * wordCount

	if strLen < totalLen {
		return []int{}
	}

	// 建立单词频率表
	wordFreq := make(map[string]int, wordCount)
	for _, word := range words {
		wordFreq[word]++
	}

	result := make([]int, 0)

	// 由于每个单词长度固定，我们可以将起始位置分成 wordLen 组
	for offset := 0; offset < wordLen; offset++ {
		// 当前窗口中的单词频率表
		currFreq := make(map[string]int)
		count := 0 // 记录当前窗口中匹配的单词数

		// 初始化第一个窗口
		for i := offset; i <= strLen-wordLen; i += wordLen {
			// 移除窗口最左边的单词
			if i >= offset+totalLen {
				leftWord := s[i-totalLen : i-totalLen+wordLen]
				if currFreq[leftWord] > 0 {
					if currFreq[leftWord] <= wordFreq[leftWord] {
						count--
					}
					currFreq[leftWord]--
				}
			}

			// 添加新单词到窗口
			word := s[i : i+wordLen]
			currFreq[word]++

			// 如果当前单词在词表中且频率不超过要求
			if freq := wordFreq[word]; freq > 0 {
				if currFreq[word] <= freq {
					count++
				}
			}

			// 检查是否找到所有单词
			if count == wordCount {
				result = append(result, i-totalLen+wordLen)
			}
		}
	}

	return result
}

// 76 最小覆盖子串 https://leetcode.cn/problems/minimum-window-substring/
// s = "ADOBECODEBANC", t = "ABC"
// minWindow 查找字符串 s 中包含字符串 t 的所有字符的最小窗口子串
// s: 源字符串
// t: 目标字符串（要查找的字符集合）
// 返回: 最小窗口子串，如果不存在则返回空字符串
func minWindow(s string, t string) string {
	// 处理边界情况：如果源字符串或目标字符串为空，返回空字符串
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// 创建数组记录目标字符串中每个字符的出现次数
	// 使用 ASCII 码作为索引，数组大小为 128
	countT := [128]int{}
	for _, char := range t {
		countT[char]++
	}

	// 初始化滑动窗口的参数
	left, right := 0, 0        // 窗口的左右边界
	minLen := len(s) + 1       // 最小窗口长度（初始化为一个不可能的大值）
	start := 0                 // 最小窗口的起始位置
	required := len(t)         // 还需要匹配的字符数量
	windowCounts := [128]int{} // 当前窗口中每个字符的出现次数

	// 移动右边界，扩展窗口
	for right < len(s) {
		char := s[right]
		windowCounts[char]++

		// 如果当前字符是目标字符串中的字符，且数量未超过需求，减少待匹配数
		if windowCounts[char] <= countT[char] {
			required--
		}

		// 当所有字符都匹配后，尝试收缩左边界
		for left <= right && required == 0 {
			// 更新最小窗口信息
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			// 移除左边界字符
			leftChar := s[left]
			windowCounts[leftChar]--
			// 如果移除的是必需字符，增加待匹配数
			if windowCounts[leftChar] < countT[leftChar] {
				required++
			}
			left++
		}
		right++
	}

	// 如果未找到符合条件的窗口，返回空字符串
	if minLen == len(s)+1 {
		return ""
	}
	// 返回最小窗口子串
	return s[start : start+minLen]
}

// 93 复原IP地址 https://leetcode.cn/problems/restore-ip-addresses/description/
// Input: "25525511135"
// Output: ["255.255.11.135", "255.255.111.35"]
// 处理每一段的有效性
func isValid(part string) bool {
	if len(part) == 0 || len(part) > 3 {
		return false
	}
	if len(part) > 1 && part[0] == '0' { // 前导零情况
		return false
	}
	val, err := strconv.Atoi(part)
	return err == nil && val >= 0 && val <= 255
}

func restoreIpAddresses(s string) (res []string) {
	// 定义构建 IP 地址的递归函数
	var construct func(k, i int, prev []byte)
	construct = func(k, i int, prev []byte) {
		// 超过 4 段或当前已处理字符串结束，直接返回
		if k > 4 {
			return
		}
		// 如果已遍历完整个字符串，且正好 4 段，存储结果
		if k == 4 && i == len(s) {
			res = append(res, string(prev))
			return
		}

		// 核心循环，处理每一段
		for j := 1; j <= 3; j++ {
			if i+j > len(s) { // 如果超出字符串长度，则结束循环
				break
			}
			// 取出当前段
			part := s[i : i+j]

			// 如果当前段值在 0-255 范围内，且没有前导零，继续递归
			if isValid(part) {
				prevTmp := append(append([]byte(nil), prev...), part...)
				if k < 3 { // 只有在前三段添加 '.' 分隔符
					prevTmp = append(prevTmp, '.')
				}
				construct(k+1, i+j, prevTmp)
			}
		}
	}

	// 从 0 段、0 索引开始，初始化 prev 为 nil
	construct(0, 0, nil)
	return
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring1("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	restoreIpAddresses("25525511135")
	fmt.Println([]rune{'z'})
}
