package binarytreedfs

import "github.com/hoffhannisyan/leetcode-75-go/solutions/structures"

/*
437. Path Sum III

Given the root of a binary tree and an integer targetSum, return the number of paths
where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards
(i.e., traveling only from parent nodes to child nodes).

Example 1:
	Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
	Output: 3
	Explanation: The paths that sum to 8 are shown.

Example 2:
	Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
	Output: 3

Constraints:
	- The number of nodes in the tree is in the range [0, 1000]
	- -10^9 <= Node.val <= 10^9
	- -1000 <= targetSum <= 1000
*/

func PathSum(root *structures.TreeNode, targetSum int) int {
	prefixSum := map[int64]int{0: 1}
	return pathSumDFS(root, 0, int64(targetSum), prefixSum)
}

func pathSumDFS(node *structures.TreeNode, currentSum int64, targetSum int64, prefixSum map[int64]int) int {
	if node == nil {
		return 0
	}

	currentSum += int64(node.Val)
	count := prefixSum[currentSum-targetSum]

	prefixSum[currentSum]++
	count += pathSumDFS(node.Left, currentSum, targetSum, prefixSum)
	count += pathSumDFS(node.Right, currentSum, targetSum, prefixSum)
	prefixSum[currentSum]--

	return count
}

/*
Time Complexity: O(n)
	- Visit each node exactly once
	- HashMap operations are O(1) average

Space Complexity: O(n)
	- HashMap stores at most n prefix sums
	- Recursion stack is O(h) where h is tree height
	- Worst case h = n for skewed tree
*/
