package bitmanipulation

/*
1318. Minimum Flips to Make a OR b Equal to c

Given 3 positives numbers a, b and c. Return the minimum flips required in
some bits of a and b to make (a OR b == c). (bitwise OR operation).
Flip operation consists of change any single bit 1 to 0 or change the bit
0 to 1 in their binary representation.

Example 1:
	Input: a = 2, b = 6, c = 5
	Output: 3
	Explanation: After flips a = 1, b = 4, c = 5 such that (a OR b == c)

Example 2:
	Input: a = 4, b = 2, c = 7
	Output: 1

Example 3:
	Input: a = 1, b = 2, c = 3
	Output: 0

Constraints:
	- 1 <= a <= 10^9
	- 1 <= b <= 10^9
	- 1 <= c <= 10^9
*/

func MinFlips(a int, b int, c int) int {
	flips := 0

	for a > 0 || b > 0 || c > 0 {
		aBit := a & 1
		bBit := b & 1
		cBit := c & 1

		switch cBit {
		case 1:
			if aBit == 0 && bBit == 0 {
				flips++
			}
		case 0:
			flips += aBit + bBit
		}

		a >>= 1
		b >>= 1
		c >>= 1
	}

	return flips
}

/*
Time Complexity: O(log(max(a, b, c)))
	- Process each bit position once
	- Maximum 30 bits for values up to 10^9

Space Complexity: O(1)
	- Only a few variables used regardless of input size
*/
