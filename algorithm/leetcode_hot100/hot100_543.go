package main

//二叉树的直径

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans = 1
	dfs(root)
	return ans - 1
}

var ans int

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	//当前节点左子树节点个数
	left := dfs(root.Left)
	//当前节点右子树节点个数
	right := dfs(root.Right)
	ans = max(ans, left+right)
	//返回子树深度
	return max(left, right) + 1
}

func main() {

}
