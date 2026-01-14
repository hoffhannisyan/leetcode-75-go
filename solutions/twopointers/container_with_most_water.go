package twopointers

/*
11. Container With Most Water

You are given an integer array height of length n. There are n vertical lines
drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the
container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:
	Input: height = [1,8,6,2,5,4,8,3,7]
	Output: 49
	Explanation: The above vertical lines are represented by array
	[1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section)
	the container can contain is 49.

Example 2:
	Input: height = [1,1]
	Output: 1

Constraints:
	- n == height.length
	- 2 <= n <= 10^5
	- 0 <= height[i] <= 10^4
*/

func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		width := right - left
		area := width * min(height[left], height[right])

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

/*
Time Complexity: O(n)
	- Single pass through the array with two pointers

Space Complexity: O(1)
	- Only using pointers and variables
*/
