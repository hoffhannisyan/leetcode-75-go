package arraystring

/*
238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal
to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:
	Input: nums = [1,2,3,4]
	Output: [24,12,8,6]

Example 2:
	Input: nums = [-1,1,0,-3,3]
	Output: [0,0,9,0,0]

Constraints:
	- 2 <= nums.length <= 10^5
	- -30 <= nums[i] <= 30
	- The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.
*/

func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	suffix := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

/*
Time Complexity: O(n)
	- Two passes through the array

Space Complexity: O(1)
	- Output array does not count as extra space
	- Only using one variable for suffix product
*/
