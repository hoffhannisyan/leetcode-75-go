package binarytreedfs

import "github.com/hoffhannisyan/leetcode-75-go/solutions/structures"

/*
1372. Longest ZigZag Path in a Binary Tree

You are given the root of a binary tree.

A ZigZag path for a binary tree is defined as follow:
	- Choose any node in the binary tree and a direction (right or left).
	- If the current direction is right, move to the right child of the current node;
	  otherwise, move to the left child.
	- Change the direction from right to left or from left to right.
	- Repeat the second and third steps until you can't move in the tree.

Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest ZigZag path contained in that tree.

Example 1:
	Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
	Output: 3
	Explanation: Longest ZigZag path in blue nodes (right -> left -> right).

Example 2:
	Input: root = [1,1,1,null,1,null,null,1,1,null,1]
	Output: 4
	Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).

Example 3:
	Input: root = [1]
	Output: 0

Constraints:
	- The number of nodes in the tree is in the range [1, 5 * 10^4]
	- 1 <= Node.val <= 100
*/

func LongestZigZag(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}
	maxLen := 0
	longestZigZagDFS(root, 0, 0, &maxLen)
	return maxLen
}

func longestZigZagDFS(node *structures.TreeNode, leftLen int, rightLen int, maxLen *int) {
	if node == nil {
		return
	}

	*maxLen = max(*maxLen, max(leftLen, rightLen))

	longestZigZagDFS(node.Left, rightLen+1, 0, maxLen)
	longestZigZagDFS(node.Right, 0, leftLen+1, maxLen)
}

/*
Time Complexity: O(n)
	- Visit each node exactly once
	- Constant work per node

Space Complexity: O(h)
	- h is the height of the tree
	- Recursion stack depth
	- Worst case O(n) for skewed tree, O(log n) for balanced tree
*/
