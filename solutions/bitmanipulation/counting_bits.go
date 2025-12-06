package bitmanipulation

/*
338. Counting Bits

Given an integer n, return an array ans of length n + 1 such that for each i
(0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

Example 1:
	Input: n = 2
	Output: [0,1,1]
	Explanation:
	0 --> 0
	1 --> 1
	2 --> 10

Example 2:
	Input: n = 5
	Output: [0,1,1,2,1,2]
	Explanation:
	0 --> 0
	1 --> 1
	2 --> 10
	3 --> 11
	4 --> 100
	5 --> 101

Constraints:
	- 0 <= n <= 10^5
*/

func CountBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		half := i >> 1
		lsb := i & 1
		ans[i] = ans[half] + lsb
	}

	return ans
}

/*
Time Complexity: O(n)
	- Single pass through numbers 1 to n
	- Each computation is O(1) using previously computed values

Space Complexity: O(n)
	- Output array of size n + 1
	- No additional space beyond the required output
*/
