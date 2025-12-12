package binarytreedfs

/*
104. Maximum Depth of Binary Tree

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path
from the root node down to the farthest leaf node.

Example 1:

	Input: root = [3,9,20,null,null,15,7]
	Output: 3

Example 2:

	Input: root = [1,null,2]
	Output: 2

Constraints:
	- The number of nodes in the tree is in the range [0, 10^4].
	- -100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	return 1 + max(left, right)
}

/*
Time Complexity: O(n)
	- Each node is visited once.

Space Complexity: O(h)
	- h is the height of the tree due to recursion stack.
*/
