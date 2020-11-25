package main

import (
	"fmt"
	"math"
)

//leetcode

//status !!!
func massage(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	for i := range dp {
		//0 接 1 没接
		dp[i] = make([]int, 2)
	}
	dp[0][0] = nums[0]
	dp[0][1] = 0

	max := func(a, b int) int {
		return int(math.Max(float64(a), float64(b)))
	}

	for i := 1; i < len(nums); i++ {
		dp[i][0] = dp[i-1][1] + nums[i]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0])
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

//节省空间
func massageAdv(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	max := func(a, b int) int {
		return int(math.Max(float64(a), float64(b)))
	}
	rest, work := 0, nums[0]
	for i := 1; i < n; i++ {
		work, rest = rest+nums[i], max(work, rest)
	}
	return max(work, rest)
}

func main() {
	fmt.Println(massage([]int{1, 2, 3, 1}))
	fmt.Println(massageAdv([]int{2, 1, 4, 5, 3, 1, 1, 3}))
}
