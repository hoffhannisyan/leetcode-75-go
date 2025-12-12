package binarytreedfs

/*
1448. Count Good Nodes in Binary Tree

Given a binary tree root, a node X in the tree is named good if in the path
from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.

Example 1:

	Input: root = [3,1,4,3,null,1,5]
	Output: 4

Example 2:

	Input: root = [3,3,null,4,2]
	Output: 3

Example 3:

	Input: root = [1]
	Output: 1

Constraints:
	- The number of nodes in the binary tree is in the range [1, 10^5].
	- Each node's value is between [-10^4, 10^4].
*/

func GoodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxSoFar {
		count = 1
		maxSoFar = node.Val
	}

	count += dfs(node.Left, maxSoFar)
	count += dfs(node.Right, maxSoFar)

	return count
}

/*
Time Complexity: O(n)
	- Each node is visited once.

Space Complexity: O(h)
	- Recursion stack, where h is the height of the tree.
*/
