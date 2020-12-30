package main

//leetcode hot100 合并二叉树

/**
思路：1、左右节点同时为null，返回空节点
2、只有一个节点为null，返回对应的节点
3、都不为空，合并
4、返回上一级
*/
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t2.Right, t2.Right)
	return t1
}

func main() {

}
