package binarytreedfs

/*
236. Lowest Common Ancestor of a Binary Tree

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined
between two nodes p and q as the lowest node in T that has both p and q as descendants
(where we allow a node to be a descendant of itself)."

Example 1:
	Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	Output: 3
	Explanation: The LCA of nodes 5 and 1 is 3.

Example 2:
	Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
	Output: 5
	Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself
	according to the LCA definition.

Example 3:
	Input: root = [1,2], p = 1, q = 2
	Output: 1

Constraints:
	- The number of nodes in the tree is in the range [2, 10^5]
	- -10^9 <= Node.val <= 10^9
	- All Node.val are unique
	- p != q
	- p and q will exist in the tree
*/

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

/*
Time Complexity: O(n)
	- Visit each node at most once
	- Early termination when target found

Space Complexity: O(h)
	- h is the height of the tree
	- Recursion stack depth
	- Worst case O(n) for skewed tree, O(log n) for balanced tree
*/
