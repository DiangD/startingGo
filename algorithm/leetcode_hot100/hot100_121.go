package main

import "math"

//最low的解法，暴力循环
func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if max < prices[j]-prices[i] {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

//优化 一次循环
func maxProfitAdv(prices []int) int {
	min, max := math.MaxInt64, 0
	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > max {
			max = price - min
		}
	}
	return max
}

//待续：dp 等想明白在写吧

func main() {

}
