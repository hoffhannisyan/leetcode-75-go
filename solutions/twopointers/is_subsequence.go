package twopointers

/*
392. Is Subsequence

Given two strings s and t, return true if s is a subsequence of t, or false
otherwise.

A subsequence of a string is a new string that is formed from the original
string by deleting some (can be none) of the characters without disturbing
the relative positions of the remaining characters. (i.e., "ace" is a
subsequence of "abcde" while "aec" is not).

Example 1:
	Input: s = "abc", t = "ahbgdc"
	Output: true

Example 2:
	Input: s = "axc", t = "ahbgdc"
	Output: false

Constraints:
	- 0 <= s.length <= 100
	- 0 <= t.length <= 10^4
	- s and t consist only of lowercase English letters.
*/

func IsSubsequence(s string, t string) bool {
	sPointer := 0
	tPointer := 0

	for sPointer < len(s) && tPointer < len(t) {
		if s[sPointer] == t[tPointer] {
			sPointer++
		}
		tPointer++
	}

	return sPointer == len(s)
}

/*
Time Complexity: O(n)
	- Single pass through string t where n is len(t)

Space Complexity: O(1)
	- Only two pointer variables used
*/
