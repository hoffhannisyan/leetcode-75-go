package binarytreebfs

/*
1161. Maximum Level Sum of a Binary Tree

Given the root of a binary tree, the level of its root is 1, the level of its
children is 2, and so on.

Return the smallest level x such that the sum of all the values of nodes at
level x is maximal.

Example 1:
	Input: root = [1,7,0,7,-8,null,null]
	Output: 2
	Explanation:
	Level 1 sum = 1.
	Level 2 sum = 7 + 0 = 7.
	Level 3 sum = 7 + -8 = -1.
	So we return the level with the maximum sum which is level 2.

Example 2:
	Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
	Output: 2

Constraints:
	- The number of nodes in the tree is in the range [1, 10^4].
	- -10^5 <= Node.val <= 10^5
*/

func MaxLevelSum(root *TreeNode) int {
	maxSum := root.Val
	maxLevel := 1
	currentLevel := 1
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if levelSum > maxSum {
			maxSum = levelSum
			maxLevel = currentLevel
		}
		currentLevel++
	}

	return maxLevel
}

/*
Time Complexity: O(n)
	- Visit each node exactly once

Space Complexity: O(w)
	- w is the maximum width of the tree (queue size)
	- Worst case O(n) for a complete binary tree
*/
