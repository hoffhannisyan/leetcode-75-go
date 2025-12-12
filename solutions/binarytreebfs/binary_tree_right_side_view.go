package binarytreebfs

/*
199. Binary Tree Right Side View

Given the root of a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Example 1:

	Input: root = [1,2,3,null,5,null,4]
	Output: [1,3,4]

Example 2:

	Input: root = [1,2,3,4,null,null,null,5]
	Output: [1,3,4,5]

Example 3:

	Input: root = [1,null,3]
	Output: [1,3]

Example 4:

	Input: root = []
	Output: []

Constraints:
	- The number of nodes in the tree is in the range [0, 100].
	- -100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		result = append(result, queue[levelSize-1].Val)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

/*
Time Complexity: O(n)
	- Each node is visited once.

Space Complexity: O(w)
	- w is the maximum width of the tree (queue size).
*/
