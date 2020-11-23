package main

import "fmt"

//初识动态规划
//首先，动态规划问题的一般形式就是求最值。动态规划其实是运筹学的一种最优化方法，只不过在计算机问题上应用比较多，比如说让你求最长递增子序列呀，最小编辑距离呀等等。

/**
1.求解动态规划的核心问题是穷举
2.存在「重叠子问题」
3.具备「最优子结构」
4.正确的「状态转移方程」

明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 dp 数组/函数的含义。
*/

//fib 暴力穷举 时间复杂度o(2^n),大量的重复计算,这就是dp的特性：重叠子问题
func fib(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

//fibAdvance 使用切片记录计算过的节点 时间复杂度o(n)=>子问题的数目
//同上自顶向下的思想
func fibAdvance(N int) int {
	if N < 1 {
		return 0
	}
	memo := make([]int, N+1)
	return helper(memo, N)
}

func helper(memo []int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}

//dp的思想是自底向上解决问题，从N=1开始。
//找到状态转移方程=>状态压缩
func fibDP(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	prev, curr := 1, 1
	for i := 3; i <= N; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

func main() {
	fmt.Println(fib(18))
	fmt.Println(fibAdvance(18))
	fmt.Println(fibDP(18))
}
