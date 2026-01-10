package stack

import "strings"

/*
394. Decode String

Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square
brackets is being repeated exactly k times. Note that k is guaranteed to be a positive
integer.

You may assume that the input string is always valid; there are no extra white spaces,
square brackets are well-formed, etc. Furthermore, you may assume that the original
data does not contain any digits and that digits are only for those repeat numbers, k.
For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 10^5.

Example 1:
	Input: s = "3[a]2[bc]"
	Output: "aaabcbc"

Example 2:
	Input: s = "3[a2[c]]"
	Output: "accaccacc"

Example 3:
	Input: s = "2[abc]3[cd]ef"
	Output: "abcabccdcdcdef"

Constraints:
	- 1 <= s.length <= 30
	- s consists of lowercase English letters, digits, and square brackets '[]'
	- s is guaranteed to be a valid input
	- All the integers in s are in the range [1, 300]
*/

type frame struct {
	str string
	num int
}

func DecodeString(s string) string {
	stack := []frame{}
	var current strings.Builder
	num := 0

	for _, char := range s {
		switch {
		case char >= '0' && char <= '9':
			num = num*10 + int(char-'0')
		case char == '[':
			stack = append(stack, frame{str: current.String(), num: num})
			current.Reset()
			num = 0
		case char == ']':
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			repeated := strings.Repeat(current.String(), top.num)
			current.Reset()
			current.WriteString(top.str)
			current.WriteString(repeated)
		default:
			current.WriteRune(char)
		}
	}

	return current.String()
}

/*
Time Complexity: O(n)
	- n is the length of the decoded output
	- Each character in output is written once
	- strings.Repeat is O(length of repeated string)

Space Complexity: O(d + n)
	- d is maximum nesting depth for the stack
	- n is the length of decoded output stored in Builder
*/
