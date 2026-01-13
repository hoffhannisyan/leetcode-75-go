package linkedlist

/*
2095. Delete the Middle Node of a Linked List

You are given the head of a linked list. Delete the middle node, and return
the head of the modified linked list.

The middle node of a linked list of size n is the ⌊n / 2⌋th node from the
start using 0-based indexing, where ⌊x⌋ denotes the largest integer less
than or equal to x.

For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.

Example 1:
	Input: head = [1,3,4,7,1,2,6]
	Output: [1,3,4,1,2,6]
	Explanation:
	Since n = 7, node 3 with value 7 is the middle node.
	We return the new list after removing this node.

Example 2:
	Input: head = [1,2,3,4]
	Output: [1,2,4]
	Explanation:
	For n = 4, node 2 with value 3 is the middle node.

Example 3:
	Input: head = [2,1]
	Output: [2]
	Explanation:
	For n = 2, node 1 with value 1 is the middle node.
	Node 0 with value 2 is the only node remaining after removing node 1.

Constraints:
	- The number of nodes in the list is in the range [1, 10^5].
	- 1 <= Node.val <= 10^5
*/

func DeleteMiddle(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

/*
Time Complexity: O(n)
	- Single pass through the list with fast pointer

Space Complexity: O(1)
	- Only using two pointers and one dummy node
*/
