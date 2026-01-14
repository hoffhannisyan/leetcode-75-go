package twopointers

import "sort"

/*
1679. Max Number of K-Sum Pairs

You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k
and remove them from the array.

Return the maximum number of operations you can perform on the array.

Example 1:
	Input: nums = [1,2,3,4], k = 5
	Output: 2
	Explanation: Starting with nums = [1,2,3,4]:
	- Remove numbers 1 and 4, then nums = [2,3]
	- Remove numbers 2 and 3, then nums = []
	There are no more pairs that sum up to 5, hence a total of 2 operations.

Example 2:
	Input: nums = [3,1,3,4,3], k = 6
	Output: 1
	Explanation: Starting with nums = [3,1,3,4,3]:
	- Remove the first two 3's, then nums = [1,4,3]
	There are no more pairs that sum up to 6, hence a total of 1 operation.

Constraints:
	- 1 <= nums.length <= 10^5
	- 1 <= nums[i] <= 10^9
	- 1 <= k <= 10^9
*/

func MaxOperations(nums []int, k int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	count := 0

	for left < right {
		sum := nums[left] + nums[right]

		if sum == k {
			count++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}

	return count
}

/*
Time Complexity: O(n log n)
	- Sorting takes O(n log n)
	- Two pointer traversal takes O(n)

Space Complexity: O(1)
	- Only using pointers and variables
	- Sorting is in-place
*/
