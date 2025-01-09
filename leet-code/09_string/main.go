package main

import "fmt"

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

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))

}
