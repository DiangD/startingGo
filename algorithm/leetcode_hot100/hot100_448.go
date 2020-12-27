package main

//leetcode hot100 448 找到所有数组中消失的数字

//hash解法 O(n)
func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]int)
	var res []int
	for _, num := range nums {
		m[num] = 1
	}

	for i := 1; i <= len(nums); i++ {
		if m[i] != 1 {
			res = append(res, i)
		}
	}
	return res
}

// 标志法
func findDisappearedNumbersAdv(nums []int) []int {
	tmp := make([]int, len(nums))
	var res []int
	for _, num := range nums {
		tmp[num] = -1
	}
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {

}
