package main

import "fmt"

/*
剑指offer 07 重建二叉树
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
A:利用前序遍历的特点，preorder[0]==root,在中序遍历里找到root，左边为左子树，右边为右子树。
利用这个原理递归
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	for index, v := range inorder {
		if v == preorder[0] {
			return &TreeNode{
				Val: v,
				//左子树的root，左子树
				Left: buildTree(preorder[1:index+1], inorder[:index]),
				//右子树的root，右子树
				Right: buildTree(preorder[index+1:], inorder[index+1:]),
			}
		}
	}
	return nil
}
func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
