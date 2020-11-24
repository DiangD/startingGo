package main

import (
	"fmt"
	"math"
)

//经典leetcode题

//思路：找出状态转移方程 => 暴力破解 => 优化子结构 => 自底向上dp

//暴力破解
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt64
	//计算子结构
	for _, coin := range coins {
		sub := coinChange(coins, amount-coin)
		if sub == -1 {
			continue
		}
		res = int(math.Min(float64(res), float64(1+sub)))
	}
	if res == math.MaxInt64 {
		return -1
	}
	return res
}

//使用备忘录
func coinChangeAdvance(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	memo := make([]int, amount+1)
	return coinHelper(memo, coins, amount)
}

func coinHelper(memo []int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	//先查找备忘录
	if memo[amount] != 0 {
		return memo[amount]
	}
	res := math.MaxInt64
	//计算子结构
	for _, coin := range coins {
		sub := coinHelper(memo, coins, amount-coin)
		if sub == -1 {
			continue
		}
		res = int(math.Min(float64(res), float64(1+sub)))
	}
	if res == math.MaxInt64 {
		memo[amount] = -1
	}
	memo[amount] = res
	return memo[amount]
}

func coinsDP(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	//建立dp数组
	dp := make([]int, amount+1)
	//初始化为amount
	for i, _ := range dp {
		dp[i] = amount
	}
	dp[0] = 0
	//自底向上dp
	for i, _ := range dp {
		for _, coin := range coins {
			//无解跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coin])))
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
	fmt.Println(coinChangeAdvance(coins, amount))
	fmt.Println(coinsDP(coins, amount))
}
