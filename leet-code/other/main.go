package main

// 前缀和
// https://leetcode.cn/problems/count-number-of-nice-subarrays/description/
func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]%2
	}
	ans := 0
	count := make([]int, n+1)
	count[sum[0]]++
	for i := 1; i <= n; i++ {
		if sum[i]-k >= 0 {
			ans += count[sum[i]-k]
		}
		count[sum[i]]++
	}
	return ans
}

func main() {}
