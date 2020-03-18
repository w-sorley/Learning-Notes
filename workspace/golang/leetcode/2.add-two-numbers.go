/*
 * @Author: wang shouli
 * @LastEditors: wang shouli
 */
/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (32.03%)
 * Likes:    6421
 * Dislikes: 1680
 * Total Accepted:    1.1M
 * Total Submissions: 3.4M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Example:
 *
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var header *ListNode
	now := header
	var forward int
	for l1 != nil || l2 != nil || forward != 0 {
		var val1 int
		if l1 == nil {
			val1 = 0
		} else {
			val1 = l1.Val
			l1 = l1.Next
		}

		var val2 int
		if l2 == nil {
			val2 = 0
			l2 = nil
		} else {
			val2 = l2.Val
			l2 = l2.Next
		}

		temp := &ListNode{
			Next: nil,
			Val:  (val1 + val2 + forward) % 10,
		}
		forward = (val1 + val2 + forward) / 10

		l1 = l1
		l2 = l2
		if header == nil {
			header = temp
			now = temp
			continue
		}
		now.Next = temp
		now = temp
	}
	return header
}

// @lc code=end

