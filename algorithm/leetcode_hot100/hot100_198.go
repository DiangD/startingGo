package main

//leetcode hot100 198 打家劫舍

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func main() {

}
